package cat

import (
	"encoding/binary"
	"errors"
	"io"
	. "os"
	"syscall"
	"time"
)

func cat_new_mids() (floor uint64, ceiling uint64, tsh uint64, err error) {
	tsh = uint64(time.Now().Unix() / 3600)
	file, err := OpenFile(TEMPFILE, O_CREATE|O_RDWR, 0664)
	if err != nil {
		return 0, 0, tsh, errors.New("Unable to open temp file")
	}
	defer func() {
		file.Sync()
		file.Close()
		syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	} ()
	share := make([]byte, 16)
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	n, err := file.Read(share)
	if err != nil && err != io.EOF {
		return 0, 0, tsh, errors.New("Unable to read temp file")
	}
	if n == 16 {
		var f uint64 = binary.BigEndian.Uint64(share[:8])
		var l uint64 = binary.BigEndian.Uint64(share[8:])
		if tsh > l {
			floor = 1
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
		floor = 1
		ceiling = 100000
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(ceiling))
		n, err = file.WriteAt(buf, 0)
		if err != nil {
			return 0, 0, tsh, errors.New("Unable to write index to temp file")
		}
		binary.BigEndian.PutUint64(buf, uint64(tsh))
		n, err = file.WriteAt(buf, 8)
		if err != nil {
			return 0, 0, tsh, errors.New("Unable to write tsh to temp file")
		}
	} else {
		return 0, 0, tsh, errors.New("Temp file is herpahs corrupted")
	}
	return floor, ceiling, tsh, nil
}
