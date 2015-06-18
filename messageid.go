package cat

type MessageIdFactory interface {
	Next() MessageId
}

type message_id_factory struct {
	index int
	index_l chan int
	next_ids chan int
}

func NewMessageIdFactory() MessageIdFactory {
	return &message_id_factory{
		0,
		make(chan int, 1),
		make(chan int, 1<<10),
	}
}

var MESSAGE_ID_FACTORY MessageIdFactory = NewMessageIdFactory()

func (f *message_id_factory) Next() MessageId {
	select {
	case nextid := <-f.next_ids:
		next := NewMessageId()
		next.SetIndex(nextid)
		return next
	default:
		go f.generate()
		return f.Next()
	}
}

func (f *message_id_factory) generate() {
	f.index_l <- 0
	for i := 0; i< 1<<10; i++ {
		f.next_ids <- f.index
		f.index++
	}
	<-f.index_l
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
