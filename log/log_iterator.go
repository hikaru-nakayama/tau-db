package log

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type LogIterator struct {
	fm              *file.FileMgr
	blk             *file.BlockId
	p               *file.Page
	currentPosition int
	boundary        int // The offset of the most recently added record.
}

func NewLogIterator(fm *file.FileMgr, blk *file.BlockId) *LogIterator {
	p := file.NewPage(fm.BlockSize())
	li := &LogIterator{
		fm:  fm,
		blk: blk,
		p:   p,
	}
	li.moveToBlock(blk)
	return li
}

func (li *LogIterator) Next() []byte {
	if li.currentPosition == li.fm.BlockSize() {
		blk := file.NewBlockId(li.blk.Filename(), li.blk.Number()-1)
		li.blk = blk
		li.moveToBlock(blk)
	}
	rec := li.p.GetBytes(li.currentPosition)
	li.currentPosition += 4 + len(rec)
	return rec
}

func (li *LogIterator) HasNext() bool {
	return li.currentPosition < li.fm.BlockSize() || li.blk.Number() > 0
}



func (li *LogIterator) moveToBlock(blk *file.BlockId) {
	li.fm.Read(blk, li.p)
	li.boundary = li.p.GetInt(0)
	li.currentPosition = li.boundary
}
