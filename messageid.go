package cat

type MessageIdFactory interface {
	Next() MessageId
}

type message_id_factory struct {
	index int
	index_l chan int
}

func NewMessageIdFactory() MessageIdFactory {
	return &message_id_factory{
		0,
		make(chan int, 1),
	}
}

var MESSAGE_ID_FACTORY MessageIdFactory = NewMessageIdFactory()

func (f *message_id_factory) Next() MessageId {
	next := NewMessageId()
	f.index_l <- 0
	next.SetIndex(f.index)
	f.index++
	<-f.index_l
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
