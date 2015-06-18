package cat

var cat_lock chan int = make(chan int, 1)
var cat_initialized bool = false

//A tool instance for CAT.
//Use it to create Transaction, Event, Heartbeat, Trace...
//Every Cat instance has 1 Tree instance.
type Cat interface {
	Tree
}

type cat struct {
	Tree
}

//Cat_init_if initialize cat.go,
//which must be down before any other operations,
//for which Instance called it automatically.
func Cat_init_if() {
	cat_lock <- 0
	if !cat_initialized {
		go sender_run()
		cat_initialized = true
	}
	<-cat_lock
}

//As it's not recommended to apply thread local in go,
//apps with cat.go have to call Instance,
//keep and manage the instance returned properly.
func Instance() Cat {
	Cat_init_if()
	return &cat{
		NewTree(),
	}
}
