package buffer

import (
	"testing"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"

)

func TestBuffer(t *testing.T) {
	fm := file.NewFileMgr("test_db", 400)
	lm, err := log.NewLogMgr(fm, "test_file")
	if err != nil {
		t.Errorf("fail to New LogMgr")
	}
	bf := NewBuffer(fm, lm)
	if bf.fm != fm || bf.lm != lm {
		t.Errorf("fail to initialize")
	}
}



	
	

