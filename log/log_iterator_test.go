package log

import (
	"testing"
        "github.com/hikaru-nakayama/tau-db.git/file"
	"os"
)

func TestNewLogIterator(t *testing.T) {
	fm := file.NewFileMgr("test_db", 400)
	blk := file.NewBlockId("testfile", 2)
	iterator := NewLogIterator(fm, blk)
	if iterator.fm != fm || iterator.blk != blk {
		t.Errorf("fail to initialize")
	}
	if iterator.boundary != 0 {
		t.Errorf("expected: %d, but got: %d", 0, iterator.boundary)
	}
	os.RemoveAll("test_db")
}

