package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type SetIntRecord struct {
	LogRecord
	txnum  int
	offset int
	val    int
	blk    *file.BlockId
}

func NewSetIntRecord(p *file.Page) *SetIntRecord {
	tpos := 4
	txnum := p.GetInt(tpos)
	fpos := tpos + 4
	filename := p.GetString(fpos)
	bpos := fpos + file.MaxLength(len(filename))
	blknum := p.GetInt(bpos)
	blk := file.NewBlockId(filename, blknum)
	opos := bpos + 4
	offset := p.GetInt(opos)
	vpos := offset + 4
	val := p.GetInt(vpos)

	return &SetIntRecord{
		txnum: txnum,
		offset: offset,
		val: val,
		blk: blk
	}
}


func (sir *SetIntRecord) Op() int {
	return SETINT
}

func (sir *SetIntRecord) TxNumber() int {
	return sir.txnum
}

func (sir *SetIntRecord) Undo(tx *Transaction) {
	tx.Pin(sir.blk)
	tx.SetInt(sir.blk, sir.offset, sir.val, false)
	tx.UnPin(sir.blk)
}

func SetIntRecordWriteToLog(lm *log.LogMgr, txnum int, blk *file.BlockId, offset int, val string) (int error) {
	tpos := 4
	fpos := tpos + 4
	bpos := fpos + file.MaxLength(len(blk.Filename()))
	opos := bpos + 4
	vops := opos + 4
	reclen := vpos + 4
	rec := make([]byte, reclen)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, SETINT)
	p.SetInt(tpos, txnum)
	p.SetString(fpos, blk.Filename())
	p.SetInt(bpos, blk.Number())
	p.SetInt(opos, offset)
	p.SetInt(vpos, val)
	return lm.Append(rec)
}


