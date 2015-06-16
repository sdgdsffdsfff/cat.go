package cat

type Tree interface {
	NewTransaction(string, string) Transaction
}

type tree struct {
	stack []Transaction
	root Transaction
}

func NewTree() Tree{
	return &tree{
		make([]Transaction, 0),
		nil,
	}
}

func(this *tree) NewTransaction(t string, n string) Transaction{
	stack := this.stack
	transaction := NewTransaction(t, n, this.flush_t)
	l := len(stack)
	if l > 0 {
		parent := stack[l-1]
		parent.AddChild(transaction)
	} else {
		this.root = transaction
	}
	return nil
}

func(this *tree) flush_t(t Transaction) {
	stack := this.stack
	current := len(stack)-1
	for ;current>-1;current-- {
		if stack[current] == t {
			this.stack = stack[:current]
		}
	}
	if current == 0 {
		Tchan <- this.root
	}
}
