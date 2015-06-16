package cat

import "time"
import "bytes"
import "fmt"

type Message interface{
	SetStatus(Status)
	Add(string, string)
	GetType() string
	GetName() string
	GetStatus() string
	GetTimestamp() time.Time
	GetData() []byte
}

type message struct {
	m_type string
	m_name string
	m_status string
	m_timestamp time.Time
	m_data *bytes.Buffer
}

func NewMessage(t string, n string) Message{
	return &message{
		m_type: t,
		m_name: n,
		m_status: "unset",
		m_timestamp: time.Now(),
		m_data: nil,
	}
}

func (m *message) SetStatus(err Status) {
	if err == nil {
		m.m_status = "0"
	} else {
		m.m_status = fmt.Sprintf("%v", err)
	}
}

func (m *message) Add(key string, value string) {
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

func (m *message) GetType() string {
	return m.m_type
}

func (m *message) GetName() string {
	return m.m_name
}

func (m *message) GetStatus() string {
	return m.m_status
}

func (m *message) GetTimestamp() time.Time {
	return m.m_timestamp
}

func (m *message) GetData() []byte {
	return m.m_data.Bytes()
}
