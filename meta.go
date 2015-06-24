package cat

import "time"
import "bytes"
import "fmt"

type Meta interface {
	SetStatus(Panic)
	AddData(string, string)
	GetType() string
	GetName() string
	GetStatus() string
	GetTimestamp() time.Time
	SetData([]byte)
	GetData() []byte
}

type meta struct {
	m_type      string
	m_name      string
	m_status    string
	m_timestamp time.Time
	m_data      *bytes.Buffer
}

func NewMeta(t string, n string) Meta {
	return &meta{
		m_type:      t,
		m_name:      n,
		m_status:    "unset",
		m_timestamp: time.Now(),
		m_data:      nil,
	}
}

func (m *meta) SetStatus(err Panic) {
	if err == nil {
		m.m_status = "0"
	} else {
		m.m_status = fmt.Sprintf("%v", err)
	}
}

func (m *meta) AddData(key string, value string) {
	if m.m_data == nil {
		m.m_data = new(bytes.Buffer)
		m.m_data.WriteString(key)
		m.m_data.WriteString("=")
		m.m_data.WriteString(value)
	} else {
		m.m_data.WriteString("&")
		m.m_data.WriteString(key)
		m.m_data.WriteString("=")
		m.m_data.WriteString(value)
	}
}

func (m *meta) GetType() string {
	return m.m_type
}

func (m *meta) GetName() string {
	return m.m_name
}

func (m *meta) GetStatus() string {
	return m.m_status
}

func (m *meta) GetTimestamp() time.Time {
	return m.m_timestamp
}

func (m *meta) SetData(data []byte) {
	m.m_data = new(bytes.Buffer)
	m.m_data.Write(data)
}

func (m *meta) GetData() []byte {
	if m.m_data != nil {
		return m.m_data.Bytes()
	} else {
		return make([]byte, 0)
	}
}
