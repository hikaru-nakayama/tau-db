package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type CommitRecord struct {
	LogRecord
	txnum int
}

func NewCommitRecord(p *file.Page) *CommitRecord {
	tpos := 4
	txnum := p.GetInt(tpos)
	return &CommitRecord{
		txnum: txnum,
	}
}

func (cr *CommitRecord) Op() int {
	return COMMIT
}

func (cr *CommitRecord) TxNumber() int {
	return cr.txnum
}

func (cr *CommitRecord) Undo(tx *Transaction) {}

func CommitRecordWriteToLog(lm *log.LogMgr, txnum int) (int, error) {
	rec := make([]byte, 2*4)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, COMMIT)
	p.SetInt(4, txnum)
	return lm.Append(rec)
}
