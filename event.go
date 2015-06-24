package cat

type Event interface {
	Message
	Complete()
}

type event struct {
	Meta
	f Function
}

func NewEvent(t string, n string, f Function) Event {
	return &event{
		NewMeta(t, n),
		f,
	}
}

func (t *event) Complete() {
	if t.f != nil {
		Invoke(t.f, t)
	}
}
