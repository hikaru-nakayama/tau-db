package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type ConcurrencyMgr struct {
	locktbl LockTable
	locks   map[file.BlockId]string
}



