package cat

type Heartbeat interface {
	Message
	Set(extension_id string, detail_id string, value string)
}

type heartbeat struct {
	Meta
	status
	f Function
}

func NewHeartbeat(t string, n string, f Function) Heartbeat {
	return &heartbeat{
		NewMeta(t, n),
		status{
			"afsdfasdf",
			make([]extension, 0),
		},
		f,
	}
}

type status struct {
	Timestamp string      `xml:"timestamp,attr"`
	Extension []extension `xml:"extension"`
}

type extension struct {
	Id               string            `xml:"id,attr"`
	ExtensionDetails []extensionDetail `xml:"extenionDetail"`
}

type extensionDetail struct {
	Id    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
}

func (h heartbeat) Set(extension_id string, detail_id string, value string) {

}
