package cat

import "net"
import "fmt"

var Tchan chan Transaction = make(chan Transaction, 256)

//sender_run is internally used and only called by Cat_init_if.
//sender_run receive Transaction from a specific channel, 
//encode Transaction and start new goroutines to send Transaction to backend.
//Basically, only 1 goroutine keeps the function.
func sender_run() {
	for {
		t := <-Tchan
		fmt.Println(t.GetName())
		fmt.Println(t.GetType())
	}
}

//sender_send is internally used and only called by sender_run
func sender_send(conn net.Conn, data []byte) {
}
