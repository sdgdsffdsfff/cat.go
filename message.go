package cat

import "bytes"

type Encodable interface {
	Encode(*bytes.Buffer) Error
}

type Message interface {
	Meta
	Encodable
}
