package main

import cat "../"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)
	cat := cat.Instance()
	tr := cat.NewTransaction("TYPE", "NAME")

	e := cat.NewEvent("Review", "New")
	e.Add("id", "12345")
	e.Add("user", "john")
	e.SetStatus("0")
	e.Complete()

	err := recover()
	tr.SetStatus(err)
	tr.Complete()

	chan1 := make(chan int)
	chan1 <- 0
}
