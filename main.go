package main

import (
	"fmt"
	"github.com/hikaru-nakayama/tau-db.git/file"
)

func main() {
	blk_id := file.NewBlockId("aaa", 23)
	fmt.Println("Hello World")
	fmt.Printf("filename: %s Blknum: %d", blk_id.Filename(), blk_id.Number())
}
