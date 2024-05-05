package buffer

import (
	"errors"
	"github.com/hikaru-nakayama/tau-db.git/file"
	"github.com/hikaru-nakayama/tau-db.git/log"
	"sync"
	"time"
)

type BufferMgr struct {
	bufferpool   []*Buffer
	numAvailable int
	mu           sync.Mutex
}

const MAX_TIME = 10 * time.Second

func NewBufferMgr(fm *file.FileMgr, lm *log.LogMgr, numbuffs int) *BufferMgr {
	bufferpool := make([]*Buffer, numbuffs)
	for i := 0; i < numbuffs; i++ {
		bufferpool[i] = NewBuffer(fm, lm)
	}
	return &BufferMgr{
		bufferpool:   bufferpool,
		numAvailable: numbuffs,
	}
}

func (bm *BufferMgr) Available() int {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	return bm.numAvailable
}

func (bm *BufferMgr) FlushAll(txnum int) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	for _, buffer := range bm.bufferpool {
		if buffer.ModifyingTx() == txnum {
			buffer.Flush()
		}
	}
}

func (bm *BufferMgr) Unpin(buff *Buffer) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	buff.Unpin()
	if !buff.IsPined() {
		bm.numAvailable++
	}
}

func (bm *BufferMgr) Pin(blk *file.BlockId) (*Buffer, error) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	timestamp := time.Now()
	buff := bm.tryToPin(blk)
	for buff == nil && !bm.waitingTooLong(timestamp) {
		time.Sleep(MAX_TIME)
		buff = bm.tryToPin(blk)
	}
	if buff == nil {
		return nil, errors.New("BufferAbortException")
	}

	return buff, nil
}

func (bm *BufferMgr) waitingTooLong(start_time time.Time) bool {
	return time.Since(start_time) > MAX_TIME
}

func (bm *BufferMgr) tryToPin(blk *file.BlockId) *Buffer {
	buff := bm.findExistingBuffer(blk)
	if buff == nil {
		buff = bm.chooseUnpinnedBuffer()
		if buff == nil {
			return nil
		}
		buff.AssignToBlock(blk)
	}
	if !buff.IsPined() {
		bm.numAvailable--
	}
	buff.Pin()
	return buff
}

func (bm *BufferMgr) findExistingBuffer(blk *file.BlockId) *Buffer {
	for _, buffer := range bm.bufferpool {
		b := buffer.Block()
		if b != nil && b.Equals(blk) {
			return buffer
		}
	}
	return nil
}

func (bm *BufferMgr) chooseUnpinnedBuffer() *Buffer {
	for _, buff := range bm.bufferpool {
		if !buff.IsPined() {
			return buff
		}
	}
	return nil
}
