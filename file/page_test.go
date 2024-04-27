package file

import "testing"

func TestNewPage(t *testing.T) {
	p := NewPage(10)
	if len(p.contents().Bytes()) != 10 {
		t.Errorf("expected: %d, but got: %d", 10, len(p.bb.Bytes()))
	}
}
