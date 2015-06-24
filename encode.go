package cat

import "bytes"
import "strconv"
import "encoding/binary"
import "strings"
import "fmt"
import "time"

type Encodable interface {
	Encode(*bytes.Buffer)
}

func (t *transaction) Encode(buf *bytes.Buffer) {
	if t.children == nil || len(t.children) == 0 {
		buf.WriteString("A")
		buf.WriteString(t.start.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.GetType())
		buf.WriteString(TAB)
		buf.WriteString(t.GetName())
		buf.WriteString(TAB)
		buf.WriteString(t.GetStatus())
		buf.WriteString(TAB)
		buf.WriteString(strconv.FormatInt(int64(t.duration/1000), 10))
		buf.WriteString("us")
		buf.WriteString(TAB)
		buf.Write(t.GetData())
		buf.WriteString(TAB)
		buf.WriteString(LF)
	} else {
		buf.WriteString("t")
		buf.WriteString(t.start.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.GetType())
		buf.WriteString(TAB)
		buf.WriteString(t.GetName())
		buf.WriteString(TAB)
		buf.WriteString(LF)
		for _, child := range t.children {
			child.Encode(buf)
		}
		buf.WriteString("T")
		buf.WriteString(t.end.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.GetType())
		buf.WriteString(TAB)
		buf.WriteString(t.GetName())
		buf.WriteString(TAB)
		buf.WriteString(t.GetStatus())
		buf.WriteString(TAB)
		buf.WriteString(strconv.FormatInt(int64(t.duration/1000), 10))
		buf.WriteString("us")
		buf.WriteString(TAB)
		buf.Write(t.GetData())
		buf.WriteString(TAB)
		buf.WriteString(LF)
	}
}

func (h event) Encode(buf *bytes.Buffer) {
	buf.WriteString("E")
	buf.WriteString(h.GetTimestamp().Format("2006-01-02 15:04:05.999"))
	buf.WriteString(TAB)
	buf.WriteString(h.GetType())
	buf.WriteString(TAB)
	buf.WriteString(h.GetName())
	buf.WriteString(TAB)
	buf.WriteString(h.GetStatus())
	buf.WriteString(TAB)
	buf.Write(h.GetData())
	buf.WriteString(TAB)
	buf.WriteString(LF)
}

func (h heartbeat) Encode(buf *bytes.Buffer) {
}

func (h header) Encode(buf *bytes.Buffer) {
	buf.WriteString("PT1")
	buf.WriteString(TAB)
	buf.WriteString(h.m_domain)
	buf.WriteString(TAB)
	buf.WriteString(h.m_hostname)
	buf.WriteString(TAB)
	buf.WriteString(h.m_ipAddress)
	buf.WriteString(TAB)
	buf.WriteString("main")
	buf.WriteString(TAB)
	buf.WriteString("1")
	buf.WriteString(TAB)
	buf.WriteString("main")
	buf.WriteString(TAB)
	MESSAGE_ID_FACTORY.Next().Encode(buf)
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString("null")
	buf.WriteString(TAB)
	buf.WriteString(LF)
}

func (mid message_id) Encode(buf *bytes.Buffer) {
	buf.WriteString(mid.GetDomain())
	buf.WriteString("-")
	buf.WriteString(iptohex(mid.GetIpAddress()))
	buf.WriteString("-")
	buf.WriteString(strconv.FormatInt(time.Now().Unix()/3600, 10))
	buf.WriteString("-")
	buf.WriteString(strconv.Itoa(mid.index))
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
