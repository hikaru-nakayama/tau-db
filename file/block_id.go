package file

import (
	"fmt"
)

type BlockId struct {
	filename string
	blknum   int
}

func NewBlockId(filename string, number int) *BlockId {
	return &BlockId{filename: filename, blknum: number}
}

func (b *BlockId) Filename() string {
	return b.filename
}

func (b *BlockId) Number() int {
	return b.blknum
}

func (b *BlockId) Equals(other *BlockId) bool {
	return b.filename == other.Filename() && b.blknum == other.Number()
}

func (b *BlockId) String() string {
	return fmt.Sprintf("[file %s, block %d]", b.filename, b.blknum)
}
