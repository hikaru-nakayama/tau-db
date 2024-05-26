package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type CheckPointRecord struct {
	LogRecord
}

func NewCheckPointRecord(p *file.Page) *CheckPointRecord {
	return &CheckPointRecord{}
}

func (cr *CheckPointRecord) Op() int {
	return CHECKPOINT
}

func (cr *CheckPointRecord) TxNumber() int {
	return -1
}

func (cr *CheckPointRecord) Undo(tx *Transaction) {}

func CheckPointRecordWriteToLog(lm *log.LogMgr, txnum int) (int, error) {
	rec := make([]byte, 4)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, CHECKPOINT)
	return lm.Append(rec)
}
