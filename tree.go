package cat

type Tree interface {
	NewTransaction(string, string) Transaction
	NewEvent(string, string) Event
	NewHeartbeat(string, string) Heartbeat
}

type tree struct {
	stack []Transaction
	root  Transaction
}

func NewTree() Tree {
	return &tree{
		make([]Transaction, 0),
		nil,
	}
}

func (this *tree) NewTransaction(t string, n string) Transaction {
	stack := this.stack
	transaction := NewTransaction(t, n, this.flush_t)
	l := len(stack)
	if l > 0 {
		parent := stack[l-1]
		parent.AddChild(transaction)
	} else {
		this.root = transaction
	}
	this.stack = append(stack, transaction)
	return transaction
}

func (this *tree) NewEvent(t string, n string) Event {
	return NewEvent(t, n, this.flush)
}

func (this *tree) NewHeartbeat(t string, n string) Heartbeat {
	return NewHeartbeat(t, n, this.flush)
}

func (this *tree) flush_t(t Transaction) {
	stack := this.stack
	current := len(stack) - 1
	for ; current > -1; current-- {
		if stack[current] == t {
			this.stack = stack[:current]
			break
		}
	}
	if current == 0 {
		sender_transaction_channel <- this.root
	}
}

func (this *tree) flush(m Message) {
	stack := this.stack
	current := len(stack) - 1
	if current == -1 {
		aggregator_message_channel <- m
	} else {
		stack[current].AddChild(m)
	}
}
