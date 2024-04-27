package file

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


