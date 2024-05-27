package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/buffer"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type RecoveryMgr struct {
	lm    *log.LogMgr
	bm    *buffer.BufferMgr
	tx    *Transaction
	txnum int
}

func NewRecoveryMgr(tx *Transaction, txnum int, lm *log.LogMgr, bm *buffer.BufferMgr) *RecoveryMgr {
	return &RecoveryMgr{
		lm:    lm,
		bm:    bm,
		tx:    tx,
		txnum: txnum,
	}
}

func (rm *RecoveryMgr) Commit() error {
	rm.bm.FlushAll(rm.txnum)
	lsn, err := CommitRecordWriteToLog(rm.lm, rm.txnum)
	if err != nil {
		return err
	}
	rm.lm.Flush(lsn)
	return nil
}

func (rm *RecoveryMgr) SetInt(buff *buffer.Buffer, offset int, newval int) (int, error) {
	oldval := buff.Contents().GetInt(offset)
	blk := buff.Block()
	return SetIntRecordWriteToLog(rm.lm, rm.txnum, blk, offset, oldval)
}

func (rm *RecoveryMgr) SetString(buff *buffer.Buffer, offset int, newval string) (int, error) {
	oldval := buff.Contents().GetString(offset)
	blk := buff.Block()
	return SetStringRecordWriteToLog(rm.lm, rm.txnum, blk, offset, oldval)
}

func (rm *RecoveryMgr) doRollBack() {
	iter := rm.lm.Iterator()
	for iter.HasNext() {
		bytes := iter.Next()
		rec := CreateLogRecord(bytes)
		if rec.TxNumber() == rm.txnum {
			if rec.Op() == START {
				return
			}
			rec.Undo(rm.tx)
		}
	}
}
