package file

import (
	"os"
	"testing"
)

func TestNewFileMgr(t *testing.T) {
	dir_name := "test_db"
	
	mgr := NewFileMgr(dir_name, 100)
	_, err := os.Stat(dir_name)
	if os.IsNotExist(err) {
		t.Errorf("Not exist test dir")
	}
	if mgr.dbDirectory != dir_name {
		t.Errorf("expected: %s, but got: %s", dir_name, mgr.dbDirectory)
	}
        f, err := os.Create("./test_db/tmp.txt")
	if err != nil {
		panic("fail to create tmp file for test")
	}
	NewFileMgr(dir_name ,100)
	_, err = os.Stat(f.Name())
	if !os.IsNotExist(err) {
		t.Errorf("tmp file not deleted")
	}
	os.Remove(dir_name)
}
