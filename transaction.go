package cat

import "time"

type Transaction interface {
	Message
	AddChild(Message) Transaction
	Complete()
}

type transaction struct {
	Meta
	f        Function
	start    time.Time
	end      time.Time
	duration time.Duration
	children []Message
}

func NewTransaction(t string, n string, f Function) Transaction {
	return &transaction{
		NewMeta(t, n),
		f,
		time.Now(),
		time.Now(),
		-1,
		nil,
	}
}

func (t *transaction) AddChild(m Message) Transaction {
	if t.children == nil {
		t.children = make([]Message, 0)
	}
	t.children = append(t.children, m)
	return t
}

func (t *transaction) Complete() {
	t.end = time.Now()
	t.duration = time.Since(t.start)
	if t.f != nil {
		Invoke(t.f, t)
	}
}
