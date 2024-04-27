package file

import (
	"bytes"
	"encoding/binary"
)

type Page struct {
	bb *bytes.Buffer
}

func NewPage(size int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, size))
}

func NewPageFromByte(b []byte) *bytes.Buffer {
	return bytes.NewBuffer(b)
}

func (p *Page) GetInt(offset int) int {
	// The int type is 32 bit, and the Bigint type is 64 bit.
	data := p.bb.Bytes()[offset : offset+4]
	value := binary.BigEndian.Uint32(data)
	return int(value)
}

func (p *Page) SetInt(offset int, n int) {

	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(n))
	copy(p.bb.Bytes()[offset:], data)
}

