package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/buffer"
)

type ITransaction interface {
	Pin(blk *file.BlockId)
	SetString(blk *file.BlockId, offset int, val string, log bool)
	SetInt(blk *file.BlockId, offset int, val int, log bool)
	UnPin(blk *file.BlockId)
}

type Transaction struct {
	ITransaction
	nextTxNum int
	recoveryMgr *RecoveryMgr
	ConcurrencyMgr *ConcurrencyMgr
	fileMgr *file.FileMgr
	bufferMgr *buffer.BufferMgr
	txnum int
	bufferList BufferList
}


