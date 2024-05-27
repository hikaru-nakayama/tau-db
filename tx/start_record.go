package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type StartRecord struct {
	LogRecord
	txnum int
}

func NewStartRecord(p *file.Page) *StartRecord {
	tpos := 4
	txnum := p.GetInt(tpos)
	return &StartRecord{
		txnum: txnum,
	}
}

func (sr *StartRecord) Op() int {
	return START
}

func (sr *StartRecord) TxNumber() int {
	return sr.txnum
}

func (sr *StartRecord) Undo(tx *Transaction) {}

func StartRecordWriteToLog(lm *log.LogMgr, txnum int) (int, error) {
	rec := make([]byte, 2*4)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, START)
	p.SetInt(4, txnum)
	return lm.Append(rec)
}
