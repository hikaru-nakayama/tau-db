package tx

import (
	"github.com/hikaru-nakayama/tau-db.git/buffer"
	"github.com/hikaru-nakayama/tau-db.git/file"
)

type BufferList struct {
	buffers map[file.BlockId]*buffer.Buffer
	pins    []file.BlockId
	bm      *buffer.BufferMgr
}

func NewBufferList(bm *buffer.BufferMgr) *BufferList {
	bl := &BufferList{
		bm:      bm,
		buffers: make(map[file.BlockId]*buffer.Buffer),
		pins:    make([]file.BlockId, 0),
	}
	return bl
}

func (bl *BufferList) GetBuffer(blk *file.BlockId) *buffer.Buffer {
	return bl.buffers[*blk]
}

func (bl *BufferList) Pin(blk *file.BlockId) error {
	buff, error := bl.bm.Pin(blk)
	if error != nil {
		return error
	}
	bl.buffers[*blk] = buff
	bl.pins = append(bl.pins, *blk)

	return nil
}

func (bl *BufferList) UnPin(blk *file.BlockId) {
	buff := bl.buffers[*blk]
	bl.bm.Unpin(buff)
	bl.removeBlockFromPins(*blk)
	if !bl.containsBlockInPins(*blk) {
		delete(bl.buffers, *blk)
	}
}

func (bl *BufferList) UnPinAll() {
	for _, blk := range bl.pins {
		buff := bl.buffers[blk]
		bl.bm.Unpin(buff)
	}

	bl.buffers = make(map[file.BlockId]*buffer.Buffer)
	bl.pins = make([]file.BlockId, 0)

}

func (bl *BufferList) containsBlockInPins(blk file.BlockId) bool {
	for _, b := range bl.pins {
		if b == blk {
			return true
		}
	}
	return false
}

func (bl *BufferList) removeBlockFromPins(blk file.BlockId) {
	for i, b := range bl.pins {
		if b == blk {
			bl.pins = append(bl.pins[:i], bl.pins[i+1:]...)
			break
		}
	}
}
