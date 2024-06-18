package tx

import (
	"sync"
	"time"

	"github.com/hikaru-nakayama/tau-db.git/file"
)

const MAX_TIME = 10 * time.Second

type LockAbortException struct{}

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

func (lt *LockTable) Slock(blk *file.BlockId) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	timestamp := time.Now()

	for lt.hasXlock(blk) && !lt.waitingTooLong(timestamp) {
		lt.cond.Wait()
	}

	if lt.hasXlock(blk) {
		panic(LockAbortException{})
	}
	val := lt.getLockVal(blk)
	lt.locks[*blk] = val + 1
}

func (lt *LockTable) Xlock(blk *file.BlockId) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	timestamp := time.Now()

	for (lt.hasOtherSlocks(blk) || lt.hasXlock(blk)) && !lt.waitingTooLong(timestamp) {
		lt.cond.Wait()
	}
	if lt.hasOtherSlocks(blk) || lt.hasXlock(blk) {
		panic(LockAbortException{})
	}
	lt.locks[*blk] = -1
}

func (lt *LockTable) Unlock(blk *file.BlockId) {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	val := lt.getLockVal(blk)
	if val > 0 {
		lt.locks[*blk] = val - 1
	} else {
		delete(lt.locks, *blk)
		lt.cond.Broadcast()
	}
}

func (lt *LockTable) hasOtherSlocks(blk *file.BlockId) bool {
	return lt.getLockVal(blk) > 0
}

func (lt *LockTable) hasXlock(blk *file.BlockId) bool {
	return lt.getLockVal(blk) < 0
}

func (lt *LockTable) getLockVal(blk *file.BlockId) int {
	val, ok := lt.locks[*blk]
	if !ok {
		return 0
	}
	return val
}

func (lt *LockTable) waitingTooLong(start_time time.Time) bool {
	return time.Since(start_time) > MAX_TIME
}
