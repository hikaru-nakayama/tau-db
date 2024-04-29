package file

import (
	"testing"
	"os"
)

func TestFile(t *testing.T) {
	fm := NewFileMgr("file_test", 400)
	blk := NewBlockId("testfile", 2)
	p1 := NewPage(fm.BlockSize())
	pos1 := 88
	p1.SetString(pos1, "abcdefghijklm")
	size := p1.MaxLength(len("abcdefghijklm"))
	pos2 := pos1 + size
	p1.SetInt(pos2, 345)
	fm.Write(blk, p1)

	p2 := NewPage(fm.BlockSize())
	fm.Read(blk, p2)
	if p2.GetInt(pos2) != 345 {
		t.Errorf("expected: %d. but got: %d", 345, p2.GetInt(pos2))
	}
	if p2.GetString(pos1) != "abcdefghijklm" {
		t.Errorf("expected: %s. but got: %s", "abcdefghijklm", p2.GetString(pos1))
	}

	os.RemoveAll("file_test")
}

