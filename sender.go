package cat

import (
	"bytes"
	"github.com/phyxdown/ghost/pool"
	"time"
)

var (
	Mchan        chan Message
	MaxBatchSize int
	p            pool.Pool
)

//cat_sender_init is internally used and only called by Cat_init_if.
func cat_sender_init() {
	Mchan = make(chan Message, 1<<10)
	MaxBatchSize = 1 << 8
	p, _ = pool.NewBlockingPool(3, 3, CONN_FACTORY)
	go sender_run()
}

//sender_run call sender_collect repeatedly.
//Basically, only 1 goroutine keeps this function.
func sender_run() {
	for {
		if sender_collect() {
			time.Sleep(1 << 16 * time.Microsecond)
		}
	}
}

func sender_collect() bool {
	messages := make(chan Message, MaxBatchSize)
	var count = 0
collect:
	for count < MaxBatchSize {
		select {
		case message := <-Mchan:
			messages <- message
			count++
		default:
			break collect
		}
	}
	close(messages)
	if count > 0 {
		sender_encode(messages, count)
		return false
	} else {
		return true
	}
}

func sender_encode(messages <-chan Message, count int) {
	datas := make(chan []byte, count)
	for message := range messages {
		buf := bytes.NewBuffer([]byte{0, 0, 0, 0})
		NewHeader().Encode(buf)
		message.Encode(buf)
		load := int32tobytes(int32(buf.Len() - 4))
		data := buf.Bytes()
		data[0] = load[0]
		data[1] = load[1]
		data[2] = load[2]
		data[3] = load[3]
		datas <- data
	}
	close(datas)
	go sender_send(datas)
}

func sender_send(datas <-chan []byte) {
	conn, err := p.Get()
	if err != nil {
		return
	}
	defer conn.Close()
	for data := range datas {
		conn.Write(data)
	}
}
