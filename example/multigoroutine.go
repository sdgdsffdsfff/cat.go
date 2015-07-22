package main

import cat "../"
import "runtime"

func main(){
	runtime.GOMAXPROCS(4)

	cat.CAT_HOST = cat.FAT
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
		_write(cat)
	}
}

func _write(cat cat.Cat){
	tr := cat.NewTransaction("TYPE", "Mul")
	defer func(){
		p := recover()
		tr.SetStatus(p)
		tr.Complete()
	}()
	panic("mul error")
}
