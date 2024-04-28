package file

import (
	"os"
	"path/filepath"
	"strings"
)

type FileMgr struct {
	dbDirectory string
	blockSize   int
	openFiles   map[string]*os.File
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
