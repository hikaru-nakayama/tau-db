package file

import "testing"

func TestNewPage(t *testing.T) {
	p := NewPage(10)
	if len(p.bb.Bytes()) != 10 {
		t.Errorf("expected: %d, but got: %d", 10, len(p.bb.Bytes()))
	}
}

func TestGetInt(t *testing.T) {
	p := NewPage(100)
	tests := []struct {
		offset int
		num    int
	}{
		{10, 53100},
		{2, 4},
	}

	for _, tt := range tests {
		p.SetInt(tt.offset, tt.num)
		val := p.GetInt(tt.offset)
		if val != tt.num {
			t.Errorf("expected: %d, but got: %d", tt.num, val)
		}
	}

}
