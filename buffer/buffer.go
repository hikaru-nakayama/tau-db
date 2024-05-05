package buffer

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type Bffer struct {
	fm       *file.FileMgr
	lm       *log.LogMgr
	contents *file.Page
	blk      *file.BlockId
	pins     int
	txnum    int
	lsn      int
}

func NewBuffer(fm *file.FileMgr, lm *log.LogMgr) *Bffer {
	p := file.NewPage(fm.BlockSize())
	return &Bffer{
		fm:       fm,
		lm:       lm,
		contents: p,
		blk:      nil,
		pins:     0,
		txnum:    -1,
		lsn:      -1,
	}
}
