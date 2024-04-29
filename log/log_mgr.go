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

func NewLogMgr(fm *file.FileMgr, logfile string) (*LogMgr, error) {
	logpage := file.NewPage(fm.BlockSize())
	logsize := fm.Length(logfile)
	lm := &LogMgr{
		fm:      fm,
		logfile: logfile,
		logpage: logpage,
	}
	var currentblk *file.BlockId
	var err error
	if logsize == 0 {
		currentblk, err = lm.appendNewBlock()
		if err != nil {
			return nil, err
		}
	} else {
	        currentblk = file.NewBlockId(logfile, logsize-1)
		fm.Read(currentblk, logpage)
	}
	lm.currentblk = currentblk
	return lm, nil
}

func (lm *LogMgr) appendNewBlock() (*file.BlockId, error) {
	blk, err := lm.fm.Append(lm.logfile)
	if err != nil {
		return nil, err
	}
	lm.logpage.SetInt(0, lm.fm.BlockSize())
	lm.fm.Write(blk, lm.logpage)
	return blk, nil
}
