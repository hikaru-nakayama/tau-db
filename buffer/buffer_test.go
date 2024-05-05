package buffer

import (
	"fmt"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
	"os"
	"testing"
)

func TestBuffer(t *testing.T) {
	fm := file.NewFileMgr("test_db", 400)
	lm, err := log.NewLogMgr(fm, "test_file")
	if err != nil {
		t.Errorf("fail to New LogMgr")
	}
	bm := NewBufferMgr(fm, lm, 3)

	blk := file.NewBlockId("test_file", 1)
	buff1, err := bm.Pin(blk)

	if err != nil {
		fmt.Printf("%v", err)
	}

	p := buff1.Contents()
	n := p.GetInt(80)
	p.SetInt(80, n+1)
	buff1.SetModified(1, 0)
	fmt.Printf("The new value is %d", n+1)
	bm.Unpin(buff1)

	buff2, err := bm.Pin(file.NewBlockId("test_file", 2))
	_, err = bm.Pin(file.NewBlockId("test_file", 3))
	_, err = bm.Pin(file.NewBlockId("test_file", 4))
	if err != nil {
		fmt.Printf("%v", err)
	}

	bm.Unpin(buff2)
	buff2, err = bm.Pin(file.NewBlockId("test_file", 1))
	if err != nil {
		fmt.Printf("%v", err)
	}

	p2 := buff2.Contents()
	p2.SetInt(80, 9999)
	buff2.SetModified(1, 0)
	bm.Unpin(buff2)

	os.RemoveAll("test_db")
}

func TestBufferFile(t *testing.T) {
	fm := file.NewFileMgr("test_db", 400)
	lm, err := log.NewLogMgr(fm, "test_file")
	if err != nil {
		t.Errorf("fail to New LogMgr")
	}
	bm := NewBufferMgr(fm, lm, 3)

	blk := file.NewBlockId("test_file", 1)
	buff1, err := bm.Pin(blk)

	if err != nil {
		fmt.Printf("%v", err)
	}

	p := buff1.Contents()
	p.SetString(80, "nakayan")
	size := p.MaxLength(len("nakayan"))
	p.SetInt(80+size, 365)
	buff1.SetModified(1, 0)
	bm.Unpin(buff1)

	buff2, err := bm.Pin(file.NewBlockId("test_file", 1))
	if err != nil {
		fmt.Printf("%v", err)
	}

	p2 := buff2.Contents()
	str := p2.GetString(80)
	val := p2.GetInt(80 + size)
	buff2.SetModified(1, 0)
	bm.Unpin(buff2)

	if str != "nakayan" {
		t.Errorf("expected: %s, but got %s", "nakayan", str)
	}

	if val != 365 {
		t.Errorf("expected: %d, but got: %d", 365, val)
	}

	os.RemoveAll("test_db")
}
