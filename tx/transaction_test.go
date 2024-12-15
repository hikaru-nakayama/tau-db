package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/buffer"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
	"os"
	"testing"
)

func TestTx(t *testing.T) {
	fm := file.NewFileMgr("test_db", 1000)
	lm, err := log.NewLogMgr(fm, "log_file")
	bm := buffer.NewBufferMgr(fm, lm, 8)
	if err != nil {
		t.Errorf("fail")
	}
	tx1 := NewTransaction(fm, lm, bm)
	blk := file.NewBlockId("test_file", 1)

	tx1.Pin(blk)
	tx1.SetInt(blk, 80, 1, false)
	tx1.SetString(blk, 40, "one", false)
	tx1.Commit()

	tx2 := NewTransaction(fm, lm, bm)
	tx2.Pin(blk)
	ival := tx2.GetInt(blk, 80)
	sval := tx2.GetString(blk, 40)

	if ival != 1 {
		t.Errorf("expected: %d, but got: %d", 1, ival)
	}

	if sval != "one" {
		t.Errorf("expected: %s, but got: %s", "one", sval)
	}
	t.Logf("initial value at location 80 = %d\n", ival)
	t.Logf("initial value at location 40 = %s\n", sval)

	newival := ival + 1
	newsval := sval + "!"
	tx2.SetInt(blk, 80, newival, true)
	tx2.SetString(blk, 40, newsval, true)
	tx2.Commit()

	tx3 := NewTransaction(fm, lm, bm)
	tx3.Pin(blk)

	ival = tx3.GetInt(blk, 80)
	sval = tx3.GetString(blk, 40)

	if ival != 2 {
		t.Errorf("expected: %d, but got: %d", 1, ival)
	}

	if sval != "one!" {
		t.Errorf("expected: %s, but got: %s", "one", sval)
	}
	t.Logf("new value at location 80 = %d\n", ival)
	t.Logf("new value at location 40 = %s\n", sval)

	tx3.SetInt(blk, 80, 9999, true)
	t.Logf("pre-rollback value at location 80 = %d\n", tx3.GetInt(blk, 80))
	// tx3.RollBack()

	os.RemoveAll("test_db")

}
