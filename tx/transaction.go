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

var mu sync.Mutex

func nextTxNumber() int {
	mu.Lock()
	defer mu.Unlock()
	num := atomic.AddInt32(&nextTxNum, 1)
	return int(num)
}
