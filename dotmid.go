package cat

import "fmt"
import "syscall"
import "os"
import "time"
import "encoding/binary"

type Dotmid interface {
	Request() (floor int, ceiling int, tsh int64)
}

type dot_mid struct{}

const dotmidFilename = ".mid"
var DOT_MID Dotmid = NewDotmid()

func NewDotmid() Dotmid {
	return &dot_mid{}
}

func (d *dot_mid) Request() (floor int, ceiling int, tsh int64) {
	file, err := os.Open(dotmidFilename)
	if err != nil {
		fmt.Println(err)
	}
	share := make([]byte, 16);
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	n, err = file.Read(share)
	if err != nil {
		fmt.Println(err)
	}
	if n == 16 {
		var f int64 = binary.BigEndian.Uint64(share[:8])
		var l int64 = binary.BigEndian.Uint64(share[8:])
		tsh := time.Now().Unix()/3600
		if now > f {
			floor = 0;
			ceiling= 100000;
		} else {
			floor = f;
			ceiling = f + 100000;
		}
	} else if n == 0 {

	} else {
		fmt.Println(".mid file is probably tampered.")
	}
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	return 90, 100, 1000000;
}
