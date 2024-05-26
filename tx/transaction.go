package tx

import "github.com/hikaru-nakayama/tau-db.git/file"

type ITransaction interface {
	Pin(blk *file.BlockId)
	SetString(blk *file.BlockId, offset int, val int, log bool)
	UnPin(blk *file.BlockId)
}

type Transaction struct {
}
