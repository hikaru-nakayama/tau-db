package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/buffer"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
	"sync"
	"sync/atomic"
)

type ITransaction interface {
	Pin(blk *file.BlockId)
	SetString(blk *file.BlockId, offset int, val string, log bool)
	SetInt(blk *file.BlockId, offset int, val int, log bool)
	UnPin(blk *file.BlockId)
}

var nextTxNum int32 = 0

const END_OF_FILE int = -1

type Transaction struct {
	ITransaction
	nextTxNum      int
	recoveryMgr    *RecoveryMgr
	concurrencyMgr *ConcurrencyMgr
	fileMgr        *file.FileMgr
	bufferMgr      *buffer.BufferMgr
	txnum          int
	bufferList     *BufferList
}

func NewTransaction(fm *file.FileMgr, lm *log.LogMgr, bm *buffer.BufferMgr) *Transaction {
	txnum := nextTxNumber()

	recoveryMgr := NewRecoveryMgr(nil, txnum, lm, bm)
	concurrencyMgr := NewConcurrencyMgr()
	mybuffers := NewBufferList(bm)

	return &Transaction{
		nextTxNum:      txnum,
		fileMgr:        fm,
		bufferMgr:      bm,
		recoveryMgr:    recoveryMgr,
		concurrencyMgr: concurrencyMgr,
		bufferList:     mybuffers,
	}
}

func (tx *Transaction) Commit() {
	tx.recoveryMgr.Commit()
	tx.concurrencyMgr.Release()
	tx.bufferList.UnPinAll()
}

func (tx *Transaction) RollBack() {
	tx.recoveryMgr.RollBack()
	tx.concurrencyMgr.Release()
	tx.bufferList.UnPinAll()
}

func (tx *Transaction) Recover() {
	tx.bufferMgr.FlushAll(tx.txnum)
	tx.recoveryMgr.Recover()
}

func (tx *Transaction) Pin(blk *file.BlockId) {
	tx.bufferList.Pin(blk)
}

func (tx *Transaction) UnPin(blk *file.BlockId) {
	tx.bufferList.UnPin(blk)
}

func (tx *Transaction) GetInt(blk *file.BlockId, offset int) int {
	tx.concurrencyMgr.Slock(blk)
	buff := tx.bufferList.GetBuffer(blk)
	return buff.Contents().GetInt(offset)
}

func (tx *Transaction) GetString(blk *file.BlockId, offset int) string {
	tx.concurrencyMgr.Slock(blk)
	buff := tx.bufferList.GetBuffer(blk)
	return buff.Contents().GetString(offset)
}

func (tx *Transaction) SetInt(blk *file.BlockId, offset int, val int, okToLog bool) error {
	tx.concurrencyMgr.Xlock(blk)
	buff := tx.bufferList.GetBuffer(blk)
	lsn := -1
	var err error

	if okToLog {
		lsn, err = tx.recoveryMgr.SetInt(buff, offset, val)
		if err != nil {
			return err
		}
	}
	p := buff.Contents()
	p.SetInt(offset, val)
	buff.SetModified(tx.txnum, lsn)
	return nil
}

func (tx *Transaction) SetString(blk *file.BlockId, offset int, val string, okToLog bool) error {
	tx.concurrencyMgr.Xlock(blk)
	buff := tx.bufferList.GetBuffer(blk)
	lsn := -1
	var err error

	if okToLog {
		lsn, err = tx.recoveryMgr.SetString(buff, offset, val)
		if err != nil {
			return err
		}
	}
	p := buff.Contents()
	p.SetString(offset, val)
	buff.SetModified(tx.txnum, lsn)
	return nil
}

func (tx *Transaction) Size(filename string) int {
	dummyblk := file.NewBlockId(filename, END_OF_FILE)
	tx.concurrencyMgr.Slock(dummyblk)
	return tx.fileMgr.Length(filename)
}

func (tx *Transaction) Append(filename string) (*file.BlockId, error) {
	dummyblk := file.NewBlockId(filename, END_OF_FILE)
	tx.concurrencyMgr.Xlock(dummyblk)
	blk, err := tx.fileMgr.Append(filename)
	if err != nil {
		return nil, err
	}
	return blk, nil
}

func (tx *Transaction) BlockSize() int {
	return tx.fileMgr.BlockSize()
}

func (tx *Transaction) AvailableBuffs() int {
	return tx.bufferMgr.Available()
}

var mu sync.Mutex

func nextTxNumber() int {
	mu.Lock()
	defer mu.Unlock()
	num := atomic.AddInt32(&nextTxNum, 1)
	return int(num)
}
