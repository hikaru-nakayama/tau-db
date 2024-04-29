package log

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type LogIterator struct {
	fm              *file.FileMgr
	blk             *file.BlockId
	p               *file.Page
	currentPosition int
	boundary        int
}


