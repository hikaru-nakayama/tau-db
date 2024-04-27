package file

import (
	"bytes"
	"encoding/binary"
	"unicode/utf8"
)

type Page struct {
	bb *bytes.Buffer
}

func NewPage(size int) *Page {
	return &Page{bytes.NewBuffer(make([]byte, size))}
}

func NewPageFromByte(b []byte) *Page {
	return &Page{bytes.NewBuffer(b)}
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

func (p *Page) GetBytes(offset int) []byte {
	length := p.GetInt(offset)
	return p.bb.Bytes()[offset+4 : offset+4+length]
}

func (p *Page) SetBytes(offset int, b []byte) {
	p.SetInt(offset, len(b))
	copy(p.bb.Bytes()[offset+4:], b)
}

func (p *Page) GetString(offset int) string {
	b := p.GetBytes(offset)
	return string(b)
}

func (p *Page) SetString(offset int, s string) {
	b := []byte(s)
	p.SetBytes(offset, b)
}

func (p *Page) MaxLength(strlen int) int {
	return 4 + strlen*utf8.UTFMax
}

func (p *Page) contents() *bytes.Buffer {
	return p.bb
}

