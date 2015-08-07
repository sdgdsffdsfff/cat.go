package main

import (
	cat "../"
	"runtime"
	"time"
)

func main(){
	cat.DOMAIN = "555555"
	cat.CAT_HOST = cat.UAT
	runtime.GOMAXPROCS(4)
	go write()
	go heartbeat()
	chan1 := make(chan int)
	chan1 <- 0
}

func write(){
	cat := cat.Instance()
	for {
		_write(cat)
		time.Sleep(100 * time.Microsecond)
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

func heartbeat(){
	for {
		second := time.Now().Second()
		_heartbeat()
		println(time.Now().Format("2006-01-02 15:04:05.999"))
		time.Sleep(time.Duration(90 - second)*1000000000)
	}
}

func _heartbeat(){
	CAT := cat.Instance()
	tr := CAT.NewTransaction("System", "Status")
	defer func(){
		tr.AddData("dumpLocked", "false")
		p := recover()
		tr.SetStatus(p)
		tr.Complete()
	}()
	h := CAT.NewHeartbeat("Heartbeat", "192.168.141.131")
	h.Set("System", "Total", "0.9")
	h.Set("System", "I/O Thread", "0.1")
	h.SetStatus("0")
	h.Complete()
}
