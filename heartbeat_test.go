package cat

import "testing"
import "fmt"
import "time"

func TestStatus(t *testing.T) {
	ts := time.Now()
	h := NewHeartbeat("system", "192.168.141.131", nil)
	h.Set("FrameworkThread", "httpthread", "0.0")
	ret := h.Get()
	te := time.Since(ts)
	fmt.Println("error", string(ret))
	fmt.Println("time: ", te)
}
