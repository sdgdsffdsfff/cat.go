package cat

import (
	"syscall"
	. "os"
	"io"
	"time"
	"encoding/binary"
)

type Dotcat interface {
	Request() (floor uint64, ceiling uint64, tsh uint64)
}

type dot_cat struct{}

const dotcatFilename = ".cat"

var DOT_MID Dotcat = NewDotcat()

func NewDotcat() Dotcat {
	return &dot_cat{}
}

func (d *dot_cat) Request() (floor uint64, ceiling uint64, tsh uint64) {
	file, err := OpenFile(dotcatFilename, O_CREATE|O_RDWR, 0664)
	if err != nil {
		return 0, 0, 0
	}
	share := make([]byte, 16)
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	tsh = uint64(time.Now().Unix() / 3600)
	n, err := file.Read(share)
	if err != nil && err != io.EOF{
		return 0, 0, 0
	}
	if n == 16 {
		var f uint64 = binary.BigEndian.Uint64(share[:8])
		var l uint64 = binary.BigEndian.Uint64(share[8:])
		if tsh > l {
			floor = 0
			ceiling = 100000
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, uint64(ceiling))
			file.WriteAt(buf, 0)
			binary.BigEndian.PutUint64(buf, uint64(tsh))
			file.WriteAt(buf, 8)
		} else {
			floor = f
			ceiling = f + 100000
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, uint64(ceiling))
			file.WriteAt(buf, 0)
		}
	} else if n == 0 {
		floor = 0
		ceiling = 100000
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(ceiling))
		n, err = file.WriteAt(buf, 0)
		if err != nil {
			return 0, 0, 0
		}
		binary.BigEndian.PutUint64(buf, uint64(tsh))
		n, err = file.WriteAt(buf, 8)
		if err != nil {
			return 0, 0, 0
		}
	} else {
		//println(".cat file is probably tampered.")
		return 0, 0, 0
	}
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	return floor, ceiling, tsh
}
