package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type RollBackRecord struct {
	LogRecord
	txnum int
}

func NewRollBackRecord(p *file.Page) *RollBackRecord {
	tpos := 4
	txnum := p.GetInt(tpos)
	return &RollBackRecord{
		txnum: txnum,
	}
}

func (rb *RollBackRecord) Op() int {
	return ROLLBACK
}

func (rb *RollBackRecord) TxNumber() int {
	return rb.txnum
}

func (rb *RollBackRecord) Undo(tx *Transaction) {}

func RollBackRecordWriteToLog(lm *log.LogMgr, txnum int) (int, error) {
	rec := make([]byte, 2*4)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, ROLLBACK)
	p.SetInt(4, txnum)
	return lm.Append(rec)
}
