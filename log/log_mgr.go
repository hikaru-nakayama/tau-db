package log

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type LogMgr struct {
	fm             *file.FileMgr
	logfile        string
	logpage        *file.Page
	currentblk     *file.BlockId
	latestLSN      int
	latestSavedLSN int
}
