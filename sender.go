package cat

import "fmt"
import "time"
import "bytes"
import "net"
import "strconv"
import "encoding/binary"
import "strings"

var Mchan chan Message = make(chan Message, 1<<20)
var MaxBatchSize int = 1 << 8

//sender_run is internally used and only called by Cat_init_if.
//sender_run call sender_collect repeatedly.
//Basically, only 1 goroutine keeps this function.
func sender_run() {
	for {
		if sender_collect() {
			time.Sleep(5*time.Microsecond)
		}
	}
}

func sender_collect() bool{
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
		encode_header(buf)
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
	sender_send(datas)
}

func sender_send(datas <-chan []byte) {
	conn, err := net.Dial("tcp", "10.2.6.99:2280")
	if err != nil {
		fmt.Println(err)
	}
	for data := range datas {
		conn.Write(data)
	}
	conn.Close()
}

const m_domain string = "555554"
const m_hostname string = "DST51752"
const m_ipAddress string = "192.168.141.131"

var sender_index int = 0
var sender_l chan int = make(chan int, 1)

func message_id() []byte {
	var timestamp = time.Now().Unix() / 3600
	var index int
	sender_l <- 0
	index = sender_index
	sender_index++
	<-sender_l
	b := new(bytes.Buffer)
	b.WriteString(m_domain)
	b.WriteString("-")
	b.WriteString(iptohex(m_ipAddress))
	b.WriteString("-")
	b.WriteString(strconv.FormatInt(timestamp, 10))
	b.WriteString("-")
	b.WriteString(strconv.Itoa(index))
	return b.Bytes()
}

func encode_header(buf *bytes.Buffer) {
	buf.WriteString("PT1")
	buf.WriteString(TAB)
	buf.WriteString(m_domain)
	buf.WriteString(TAB)
	buf.WriteString(m_hostname)
	buf.WriteString(TAB)
	buf.WriteString(m_ipAddress)
	buf.WriteString(TAB)
	buf.WriteString("main")
	buf.WriteString(TAB)
	buf.WriteString("1")
	buf.WriteString(TAB)
	buf.WriteString("main")
	buf.WriteString(TAB)
	buf.Write(message_id())
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString(LF)
}

func int32tobytes(i int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}

func iptohex(ip string) string {
	var strs []string = strings.Split(ip, ".")
	for i := 0; i < 4; i++ {
		digit, _ := strconv.Atoi(strs[i])
		strs[i] = fmt.Sprintf("%x", digit)
	}
	return strings.Join(strs, "")
}
