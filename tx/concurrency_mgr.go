package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type ConcurrencyMgr struct {
	locktbl LockTable
	locks   map[file.BlockId]string
}

func NewConcurrencyMgr() *ConcurrencyMgr {
	cm := &ConcurrencyMgr{
		locktbl: *NewLockTable(),
		locks:   make(map[file.BlockId]string),
	}
	return cm
}

func (cm *ConcurrencyMgr) Slock(blk *file.BlockId) {
	if _, exists := cm.locks[*blk]; !exists {
		cm.locktbl.Slock(blk)
		cm.locks[*blk] = "S"
	}
}

func (cm *ConcurrencyMgr) Xlock(blk *file.BlockId) {
	if !cm.hasXlock(blk) {
		// Slock を取ることで、他のトランザクションが Slock を取ると, hasOtherSlocks が true になり、他のトランザクションが Xlock を取れないようにする。
		cm.Slock(blk)
		cm.locktbl.Xlock(blk)
		cm.locks[*blk] = "X"
	}
}

func (cm *ConcurrencyMgr) Release() {
	for blk := range cm.locks {
		cm.locktbl.Unlock(&blk)
		delete(cm.locks, blk)
	}
}

func (cm *ConcurrencyMgr) hasXlock(blk *file.BlockId) bool {
	lockType, ok := cm.locks[*blk]

	return !!ok && lockType == "X"
}
