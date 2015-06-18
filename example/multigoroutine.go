package main

import cat "../"
import "runtime"

func main(){
	runtime.GOMAXPROCS(4)

	go write()
	go write()
	go write()
	go write()

	chan1 := make(chan int)
	chan1 <- 0
}


func write(){
	cat := cat.Instance()
	for {
		tr := cat.NewTransaction("TYPE", "Mul")
		err := recover()
		tr.SetStatus(err)
		tr.Complete()
	}
}
