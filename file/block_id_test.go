package file

import "testing"

func TestEqual(t *testing.T) {
	b1 := BlockId{filename: "test", blknum: 1}

	tests := []struct {
		filename string
		num      int
		expected bool
	}{
		{"test", 1, true},
		{"dumy", 1, false},
		{"test", 2, false},
		{"dumy", 2, false},
	}

	for _, tt := range tests {
		b2 := BlockId{filename: tt.filename, blknum: tt.num}
		if b1.Equals(&b2) != tt.expected {
			t.Errorf("Equals return value is wrong")
		}
	}
}

func TestString(t *testing.T) {
	b := BlockId{filename: "student", blknum: 10}

	if "[file student, block 10]" != b.String() {
		t.Errorf("expected=[file student, block 10]")
	}
}
