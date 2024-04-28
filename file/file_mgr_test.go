package file

import (
	"testing"
	"os"
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
	os.Remove(dir_name)
}
