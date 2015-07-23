package cat

import "fmt"

var cat_lock chan int = make(chan int, 1)
var cat_initialized bool = false

//A tool instance for CAT.
//Use it to create Transaction, Event, Heartbeat, Trace...
//Every Cat instance has 1 Tree instance.
type Cat interface {
	//Tree provides methods to create different kinds of messages
	Tree
	//Create a new simple Event without tags.
	LogEvent(t string, n string)
	//Create a new Event whose type is error and status is ERROR,
	//nil is ignored
	LogError(e error)
	//Create a new Event whose type is panic and status is ERROR,
	//nil is ignored
	LogPanic(e Panic)
}

type cat struct {
	Tree
}

//LogEvent
func (c *cat) LogEvent(t string, n string) {
	e := c.NewEvent(t, n)
	e.SetStatus("0")
	e.Complete()
}

//LogError
func (c *cat) LogError(err error) {
	if err != nil {
		e := c.NewEvent("Error", err.Error())
		e.SetStatus("ERROR")
		e.Complete()
	}
}

//LogPanic
func (c *cat) LogPanic(err Panic) {
	if err != nil {
		e := c.NewEvent("Error", fmt.Sprintf("%v", err))
		e.SetStatus("ERROR")
		e.Complete()
	}
}

//Cat_init_if initialize cat.go,
//which must be down before any other operations,
//for which Instance called it automatically.
func Cat_init_if() {
	cat_lock <- 0
	if !cat_initialized {
		cat_config_init()
		cat_sender_init()
		cat_aggregator_init()
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
