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

func (cr *RollBackRecord) Op() int {
	return ROLLBACK
}

func (cr *RollBackRecord) TxNumber() int {
	return cr.txnum
}

func (cr *RollBackRecord) Undo() {}

func RollBackRecordWriteToLog(lm *log.LogMgr, txnum int) (int, error) {
	rec := make([]byte, 2*4)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, ROLLBACK)
	p.SetInt(4, txnum)
	return lm.Append(rec)
}
