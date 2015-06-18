package main

import cat "../"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)

	cat := cat.Instance()
	tr := cat.NewTransaction("TYPE", "NAME")

	err := recover()
	tr.SetStatus(err)
	tr.Complete()

	chan1 := make(chan int)
	chan1 <- 0
}
