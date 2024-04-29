package log

import (
	"fmt"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"sync"
)

type LogMgr struct {
	fm             *file.FileMgr
	logfile        string
	logpage        *file.Page
	currentblk     *file.BlockId
	latestLSN      int
	latestSavedLSN int
	mu             sync.Mutex
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

func (lm *LogMgr) Flush(lsn int) {
	if lsn >= lm.latestLSN {
		lm.flush()
	}
}

func (lm *LogMgr) Append(logrec []byte) (int, error) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	boundary := lm.logpage.GetInt(0)
	recsize := len(logrec)
	byteneeded := 4 + recsize
	var err error
	if boundary-byteneeded < 4 {
		lm.flush()
		lm.currentblk, err = lm.appendNewBlock()
		if err != nil {
			return 0, fmt.Errorf("Fail to append new block")
		}
		boundary = lm.logpage.GetInt(0)
	}
	recpos := boundary - byteneeded
	lm.logpage.SetBytes(recpos, logrec)
	lm.logpage.SetInt(0, recpos)
	lm.latestLSN += 1
	return lm.latestLSN, nil
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

func (lm *LogMgr) flush() {
	lm.fm.Write(lm.currentblk, lm.logpage)
	lm.latestSavedLSN = lm.latestLSN
}
