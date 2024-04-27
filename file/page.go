package file

import (
	"bytes"
)

type Page struct {
	bb *bytes.Buffer
}

func NewPage(size int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, size))
}
