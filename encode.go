package cat

import "bytes"
import "strconv"

type Encodable interface {
	Encode(*bytes.Buffer) Error
}

func (t *transaction) Encode(buf *bytes.Buffer) Error {
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
	return recover()
}
