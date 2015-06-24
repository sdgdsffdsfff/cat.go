package cat

import "time"
import "encoding/xml"

type Heartbeat interface {
	Message
	Set(extension_id string, extension_detail_id string, value string)
	Complete()
}

//refactor expected
type heartbeat struct {
	Meta
	s status
	f Function
}

func NewHeartbeat(t string, n string, f Function) Heartbeat {
	return &heartbeat{
		NewMeta(t, n),
		_status(),
		f,
	}
}

type status struct {
	Timestamp  string       `xml:"timestamp,attr"`
	Extensions []*extension `xml:"extension"`
}

func _status() status {
	return status{
		time.Now().Format("2006-01-02 15:04:05.999"),
		make([]*extension, 0),
	}
}

type extension struct {
	Id               string            `xml:"id,attr"`
	ExtensionDetails []extensionDetail `xml:"extenionDetail"`
}

func newextension(id string) *extension {
	return &extension{
		id,
		make([]extensionDetail, 0),
	}
}

type extensionDetail struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

func _extensionDetail(id string, value string) extensionDetail {
	return extensionDetail{
		id,
		value,
	}
}

func (h *heartbeat) Set(extension_id string, extension_detail_id string, value string) {
	for _, e := range h.s.Extensions {
		if e.Id == extension_id {
			e.ExtensionDetails = append(e.ExtensionDetails, _extensionDetail(extension_detail_id, value))
			return
		}
	}
	e := newextension(extension_id)
	e.ExtensionDetails = append(e.ExtensionDetails, _extensionDetail(extension_detail_id, value))
	h.s.Extensions = append(h.s.Extensions, e)
	return
}

func (h *heartbeat) Complete() {
	bytes, _ := xml.MarshalIndent(h.s, "", " ")
	h.SetData(append([]byte(xml.Header), bytes...))
	if h.f != nil {
		Invoke(h.f, h)
	}
}
