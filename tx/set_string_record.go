package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
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
