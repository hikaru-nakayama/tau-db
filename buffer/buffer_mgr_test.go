package buffer

import (
	"testing"
	"fmt"
        "github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"

)

func TestBufferMgr(t *testing.T) {
        fm := file.NewFileMgr("test_db", 400)
	lm, err := log.NewLogMgr(fm, "test_file")
	if err != nil {
		t.Errorf("fail to New LogMgr")
	}
	bm := NewBufferMgr(fm, lm, 3)
	buff := make([]*Buffer, 6)
	buff[0], err = bm.Pin(file.NewBlockId("test_file", 0))
	buff[1], err = bm.Pin(file.NewBlockId("test_file", 1))
	buff[2], err = bm.Pin(file.NewBlockId("test_file", 2))
	bm.Unpin(buff[1])
	buff[1] = nil
	buff[3], err = bm.Pin(file.NewBlockId("test_file", 0))
	buff[4], err = bm.Pin(file.NewBlockId("test_file", 1))

	fmt.Printf("Available buffers: %d\n", bm.Available())
	fmt.Println("Attempting to pin block 3...")

	buff[5], err = bm.Pin(file.NewBlockId("test_file", 3))

	if err != nil {
		fmt.Printf("Exception: No available buffer\n")
	}

	bm.Unpin(buff[2])
	buff[2] = nil
	buff[5], err = bm.Pin(file.NewBlockId("test_file", 3))

	if err != nil {
		t.Errorf("Exception: No available buffer\n")
	}

	fmt.Printf("Final Buffer Allocation: ")

	for i := 0; i < len(buff); i++ {
		b := buff[i]
		if b != nil {
			fmt.Printf("buff[%d] pinned to block %d ", i, b.Block().Number())
		}
	}

}


