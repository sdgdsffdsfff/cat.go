package cat

//import "fmt"
import "syscall"
import . "os"
import "time"
import "encoding/binary"

type Dotmid interface {
	Request() (floor uint64, ceiling uint64, tsh uint64)
}

type dot_mid struct{}

const dotmidFilename = "/home/phyxdown/.mid"

var DOT_MID Dotmid = NewDotmid()

func NewDotmid() Dotmid {
	return &dot_mid{}
}

func (d *dot_mid) Request() (floor uint64, ceiling uint64, tsh uint64) {
	file, err := OpenFile(dotmidFilename, O_RDWR, ModeAppend)
	if err != nil {
		//fmt.Println("#1 ", err)
		return 0, 0, 0
	}
	share := make([]byte, 16)
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	tsh = uint64(time.Now().Unix() / 3600)
	n, err := file.Read(share)
	if err != nil {
		//fmt.Println("#2 ", n, err)
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
			//fmt.Println("#3 ", n, err)
			return 0, 0, 0
		}

		binary.BigEndian.PutUint64(buf, uint64(tsh))
		n, err = file.WriteAt(buf, 8)
		if err != nil {
			//fmt.Println("#4 ", n, err)
			return 0, 0, 0
		}
	} else {
		//fmt.Println(".mid file is probably tampered.")
		return 0, 0, 0
	}
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	//fmt.Printf("%d %d %d\n", floor, ceiling, tsh)
	return floor, ceiling, tsh
}
