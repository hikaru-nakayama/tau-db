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
