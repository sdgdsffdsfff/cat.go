package cat

import "time"
import "encoding/xml"
import "strconv"

type Heartbeat interface {
	Message
	Set(extension_id string, extension_detail_id string, value string)
	Complete()
}

//refactor expected
type heartbeat struct {
	Meta
	s *status
	f Function
}

func NewHeartbeat(t string, n string, f Function) Heartbeat {
	return &heartbeat{
		NewMeta(t, n),
		newstatus(),
		f,
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
	bytes, _ := xml.MarshalIndent(h.s, "", "\r")
	h.SetData(append([]byte(xml.Header), bytes...))
	if h.f != nil {
		Invoke(h.f, h)
	}
}

type status struct {
	Timestamp  string       `xml:"timestamp,attr"`
	Runtime    *runtime     `xml:"runtime"`
	Os         *os          `xml:"os"`
	Disk       *disk        `xml:"disk"`
	Memory     *memory      `xml:"memory"`
	Thread     *thread      `xml:"thread"`
	Message    *message     `xml:"message"`
	Extensions []*extension `xml:"extension"`
}

func newstatus() *status {
	return &status{
		time.Now().Format("2006-01-02 15:04:05.999"),
		&runtime{
			strconv.FormatInt(time.Now().UnixNano()/1000000, 10),
			strconv.FormatInt(time.Now().Unix()/60, 10),
			"4.0.30319.296",
			"deploy",
			"/home/phyxdown",
		}, &os{
			"Linux",
			"amd64",
			"3.10.0-123.el7.x86_64",
			"4",
			"0",
			"1312",
			"8589934592",
			//"0",
			//"0",
		}, &disk{
			&diskvolume{
				"/", "261621313536", "201279385600", "198622265344",
			},
		}, &memory{
			"91889664", "75637312", "8793096", &gc{"Gen 0", "1"},
		},
		&thread{
			"34", "34", &dump{},
		},
		&message{
			"1922", "1", "3258904",
		},
		make([]*extension, 0),
	}
}

type runtime struct {
	StartTime   string `xml:"start-time,attr"`
	UpTime      string `xml:"up-time,attr"`
	JavaVersion string `xml:"java-version,attr"`
	UserName    string `xml:"user-name,attr"`
	UserDir     string `xml:"user-dir"`
}

type os struct {
	Name                string `xml:"name,attr"`
	Arch                string `xml:"arch,attr"`
	Version             string `xml:"version,attr"`
	AvailableProcessors string `xml:"available-processors,attr"`
	SystemLoadAverage   string `xml:"system-load-average,attr"`
	ProcessTime         string `xml:"process-time,attr"`
	TotalPhysicalMemory string `xml:"total-physical-memory,attr"`
	//AssembliesLoad string `xml:"assemblies-load,attr"`
	//ClassLoad string `xml:"class-loaded,attr"`
}

type disk struct {
	DiskVolume *diskvolume `xml:"disk-volume"`
}

type diskvolume struct {
	Id     string `xml:"id,attr"`
	Total  string `xml:"total,attr"`
	Free   string `xml:"free,attr"`
	Usable string `xml:"usable,attr"`
}

type memory struct {
	PrivateMemorySize string `xml:"private-memory-size,attr"`
	WorkingSetSize    string `xml:"working-set-size,attr"`
	HeapTotalMemory   string `xml:"heap-total-memory,attr"`
	Gc                *gc    `xml:"gc"`
}

type gc struct {
	Name  string `xml:"name,attr"`
	Count string `xml:"count,attr"`
}

type thread struct {
	Count             string `xml:"count,attr"`
	TotalStartedCount string `xml:"total-started-count,attr"`
	Dump              *dump  `xml:"dump"`
}

type dump struct{}

type message struct {
	Produced   string `xml:"produced,attr"`
	Overflowed string `xml:"overflowed,attr"`
	Bytes      string `xml:"bytes,attr"`
}

type extension struct {
	Id               string            `xml:"id,attr"`
	ExtensionDetails []extensionDetail `xml:"extensionDetail"`
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
