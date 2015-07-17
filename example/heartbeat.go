package main

import cat "../"
import "runtime"

func main(){
	runtime.GOMAXPROCS(4)
	mycat := cat.Instance()
	_write(mycat)
	chan1 := make(chan int)
	chan1 <- 0
}

func _write(mycat cat.Cat){
	tr := mycat.NewTransaction("System", "Status")
	defer func(){
		tr.AddData("dumpLocked", "false")
		p := recover()
		tr.SetStatus(p)
		tr.Complete()
	}()

	h := mycat.NewHeartbeat("Heartbeat", "192.168.141.131")
	h.Set("System", "Total", "0.9")
	h.Set("System", "I/O Thread", "0.1")
	h.SetStatus("0")
	h.Complete()
}
