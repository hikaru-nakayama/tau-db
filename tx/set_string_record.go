package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
)

type SetStringRecord struct {
	LogRecord
	txnum  int
	offset int
	val    string
	blk    *file.BlockId
}

func NewSetStringRecord(p *file.Page) *SetStringRecord {
	tpos := 4
	txnum := p.GetInt(tpos)
	fpos := tpos + 4
	filename := p.GetString(fpos)
	bpos := fpos + p.MaxLength(len(filename))
	blknum := p.GetInt(bpos)
	blk := file.NewBlockId(filename, blknum)
	opos := bpos + 4
	offset := p.GetInt(opos)
	vpos := opos + 4
	val := p.GetString(vpos)

	return &SetStringRecord{
		txnum:  txnum,
		offset: offset,
		val:    val,
		blk:    blk,
	}
}

func (ssr *SetStringRecord) Op() int {
	return SETSTRING
}

func (ssr *SetStringRecord) TxNumber() int {
	return ssr.txnum
}

func SetStringRecordWriteToLog(lm *log.LogMgr, txnum int, blk *file.BlockId, offset int, val string) (int, error) {
	tpos := 4
	fpos := tpos + 4
	bpos := fpos + file.MaxLength(len(blk.Filename()))
	opos := bpos + 4
	vpos := opos + 4
	reclen := vpos + file.MaxLength(len(val))
	rec := make([]byte, reclen)
	p := file.NewPageFromByte(rec)
	p.SetInt(0, SETSTRING)
	p.SetInt(tpos, txnum)
	p.SetString(fpos, blk.Filename())
	p.SetInt(bpos, blk.Number())
	p.SetInt(opos, offset)
	p.SetString(vpos, val)
	return lm.Append(rec)
}
