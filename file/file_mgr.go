package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FileMgr struct {
	dbDirectory string
	blockSize   int
	openFiles   map[string]*os.File
	mu          sync.Mutex
}

func NewFileMgr(dbDirectory string, blockSize int) *FileMgr {
	_, err := os.Stat(dbDirectory)
	if os.IsNotExist(err) {
		err := os.Mkdir(dbDirectory, 0750)
		if err != nil {
			panic("Fail to create dir for file maneger")
		}
	}

	files, err := os.ReadDir(dbDirectory)
	if err != nil {
		panic("Fail to open db dir for file maneger")
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "tmp") {
			os.Remove(filepath.Join(dbDirectory, file.Name()))
		}
	}

	return &FileMgr{
		dbDirectory: dbDirectory,
		blockSize:   blockSize,
		openFiles:   make(map[string]*os.File),
	}

}

func (fmgr *FileMgr) Read(blk *BlockId, p *Page) error {
	fmgr.mu.Lock()
	defer fmgr.mu.Unlock()
	f, err := fmgr.getFile(blk.Filename())
	if err != nil {
		return fmt.Errorf("can not read block. error: %w", err)
	}
	_, err = f.Seek(int64(blk.Number()*fmgr.blockSize), 0)
	if err != nil {
		return err
	}
	_, err = f.Read(p.contents().Bytes())
	return err
}

func (fmgr *FileMgr) Write(blk *BlockId, p *Page) error {
	fmgr.mu.Lock()
	defer fmgr.mu.Unlock()
	f, err := fmgr.getFile(blk.Filename())
	if err != nil {
		return fmt.Errorf("can not read block. error: %w", err)
	}
	_, err = f.Seek(int64(blk.Number()*fmgr.blockSize), 0)
	if err != nil {
		return err
	}
	_, err = f.Write(p.contents().Bytes())
	return err
}

func (fmgr *FileMgr) getFile(filename string) (*os.File, error) {
	f, ok := fmgr.openFiles[filename]
	if !ok {
		f, err := os.Create(filepath.Join(fmgr.dbDirectory, filename))
		if err != nil {
			return nil, err
		}
		fmgr.openFiles[filename] = f
	}
	return f, nil
}
