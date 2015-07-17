package cat

type MessageIdFactory interface {
	Next() MessageId
}

type message_id_factory struct {
	index   int
	ceiling int
	index_l chan int
}

func NewMessageIdFactory() MessageIdFactory {
	return &message_id_factory{
		0,
		0,
		make(chan int, 1),
	}
}

var MESSAGE_ID_FACTORY MessageIdFactory = NewMessageIdFactory()

func (f *message_id_factory) requestForFreshIds() {
	f.index, f.ceiling = DOT_MID.Request()
}

func (f *message_id_factory) Next() MessageId {
	f.index_l <- 0
	if !(f.index < f.ceiling) {
		f.requestForFreshIds();
	}
	index := f.index
	f.index++
	<-f.index_l
	next := NewMessageId()
	next.SetIndex(index)
	return next
}

type MessageId interface {
	Encodable
	SetIndex(index int)
}

type message_id struct {
	Header
	index int
}

func NewMessageId() MessageId {
	return &message_id{
		NewHeader(),
		-1,
	}
}

func (mid *message_id) SetIndex(index int) {
	mid.index = index
}
