package cat

import "testing"
import "encoding/xml"
import "fmt"

func TestStatus(t *testing.T) {
	ed := make([]extensionDetail, 0)
	ed = append(ed, extensionDetail{
		"httpthread",
		"0.0",
	})

	e := make([]extension, 0)
	e = append(e, extension{
		"FrameworkThread",
		ed,
	})
	s := status{
		"2015-06-23 13:43:18.767",
		e,
	}
	bytes, _ := xml.Marshal(s)
	fmt.Printf("%s\n", bytes)
}
