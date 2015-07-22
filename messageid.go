package cat

type MessageIdFactory interface {
	Next() MessageId
}

type message_id_factory struct {
	index   uint64
	ceiling uint64
	tsh     uint64
	lock    chan int
}

func NewMessageIdFactory() MessageIdFactory {
	return &message_id_factory{
		0,
		0,
		0,
		make(chan int, 1),
	}
}

var MESSAGE_ID_FACTORY MessageIdFactory = NewMessageIdFactory()

func (f *message_id_factory) requestForFreshIds() {
	f.index, f.ceiling, f.tsh = DOT_MID.Request()
}

func (f *message_id_factory) Next() MessageId {
	f.lock <- 0
	if !(f.index < f.ceiling) {
		f.requestForFreshIds()
	}
	index := f.index
	tsh := f.tsh
	f.index++
	<-f.lock
	next := NewMessageId()
	next.SetIndex(index)
	next.SetTsh(tsh)
	return next
}

type MessageId interface {
	Encodable
	SetIndex(index uint64)
	SetTsh(tsh uint64)
}

type message_id struct {
	Header
	index uint64
	tsh   uint64
}

func NewMessageId() MessageId {
	return &message_id{
		NewHeader(),
		0,
		0,
	}
}

func (mid *message_id) SetIndex(index uint64) {
	mid.index = index
}

func (mid *message_id) SetTsh(tsh uint64) {
	mid.tsh = tsh
}
