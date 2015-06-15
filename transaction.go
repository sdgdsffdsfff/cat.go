package cat

import "fmt"

import "time"
import "bytes"
import "strconv"

type transaction struct {
	tree *Tree
	*meta
	start time.Time
	end time.Time
	duration time.Duration
	children []Message
}

func newtransaction(t string, n string, tree *Tree) *transaction{
	return &transaction{
		tree,
		newmeta(t, n),
		time.Now(),
		time.Now(),
		-1,
		nil,
	}
}

func (t *transaction) addChild(m Message) *transaction{
	if t.children == nil {
		t.children = make([]Message, 0)
	}
	t.children = append(t.children, m)
	return t
}

func (t *transaction) SetStatus(err interface{}) {
	if err == nil {
		t.m_status = "0"
	} else {
		t.m_status = fmt.Sprintf("%v", err)
	}
}

func (t *transaction) Add(key string, value string) {
	if t.m_data == nil {
		t.m_data = new(bytes.Buffer)
		t.m_data.WriteString(key)
		t.m_data.WriteString("=")
		t.m_data.WriteString(value)
	} else {
		t.m_data.WriteString("&")
		t.m_data.WriteString(key)
		t.m_data.WriteString("=")
		t.m_data.WriteString(value)
	}
}

func (t *transaction) Complete() {
	t.end = time.Now()
	t.duration = time.Since(t.start)
	t.tree.flush(t)
}

func writestring(buf *bytes.Buffer, s string) int {
	c, _ := buf.WriteString(s)
	return c
}

func (t *transaction) Encode(buf *bytes.Buffer) (count int, err error) {
	count = 0
	err = nil
	if t.children == nil || len(t.children) == 0 {
		buf.WriteString("A")
		buf.WriteString(t.start.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.m_type)
		buf.WriteString(TAB)
		buf.WriteString(t.m_name)
		buf.WriteString(TAB)
		buf.WriteString(t.m_status)
		buf.WriteString(TAB)
		buf.WriteString(strconv.FormatInt(int64(t.duration/1000), 10))
		buf.WriteString("us")
		buf.WriteString(TAB)
		buf.Write(t.m_data.Bytes())
		buf.WriteString(TAB)
		buf.WriteString(LF)
	} else {
		buf.WriteString("t")
		buf.WriteString(t.start.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.m_type)
		buf.WriteString(TAB)
		buf.WriteString(t.m_name)
		buf.WriteString(TAB)
		buf.WriteString(LF)
		for _, child := range t.children {
			child.Encode(buf)
		}
		buf.WriteString("T")
		buf.WriteString(t.end.Format("2006-01-02 15:04:05.999"))
		buf.WriteString(TAB)
		buf.WriteString(t.m_type)
		buf.WriteString(TAB)
		buf.WriteString(t.m_name)
		buf.WriteString(TAB)
		buf.WriteString(t.m_status)
		buf.WriteString(TAB)
		buf.WriteString(strconv.FormatInt(int64(t.duration/1000), 10))
		buf.WriteString("us")
		buf.WriteString(TAB)
		buf.Write(t.m_data.Bytes())
		buf.WriteString(TAB)
		buf.WriteString(LF)
	}
	return
}
