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

func TestLength(t *testing.T) {
	mgr := NewFileMgr("test_db", 100)
	f, err := mgr.getFile("testfile")
	if err != nil {
		t.Errorf("can not read block")
	}
	data := make([]byte, 350)
	_, err = f.Write(data)
	if err != nil {
		t.Errorf("can not write block, err %s", err)
	}
	l := mgr.Length("testfile")
	if l != 3 {
		t.Errorf("expected: 3, but got: %d", l)
	}
	os.RemoveAll("test_db")
}
