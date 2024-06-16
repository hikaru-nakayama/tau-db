package tx

import (
	"sync"
	"time"

	"github.com/hikaru-nakayama/tau-db.git/file"
)

const MAX_TIME = 10 * time.Second

type LockTable struct {
	locks map[file.BlockId]int
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewLockTable() *LockTable {
	lt := &LockTable{
		locks: make(map[file.BlockId]int),
	}
	lt.cond = sync.NewCond(&lt.mu)
	return lt
}
