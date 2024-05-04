package log

import (
	"testing"
	"log"
	"github.com/hikaru-nakayama/tau-db.git/file"
        "fmt"
	"os"
)

func TestLog(t *testing.T) {
	fm := file.NewFileMgr("test_db", 400)
	lm, err := NewLogMgr(fm, "log_file")
	if err != nil {
		log.Fatal("occour error")
	}

	createReocords(1, 35, lm)
	printLogRecords("The log file new these records: ", lm)


	os.RemoveAll("test_db")
}

func printLogRecords(msg string, lm *LogMgr) {
	fmt.Printf("%s", msg)
	iter := lm.Iterator()
	for iter.HasNext() {
		rec := iter.Next()
		p := file.NewPageFromByte(rec)
		s := p.GetString(0)
		npos := p.MaxLength(len(s))
		val := p.GetInt(npos)
		fmt.Printf("[%s, %d]",s, val)
	}
	fmt.Println("")
}

func createReocords(start int, end int, lm *LogMgr) {
   log.Print("Creating records: ")
   for i := start; i <= end; i++ {
	   str := fmt.Sprintf("record%d", i)
	   rec := createLogRecord(str, i+100, lm)
	   lsn, err := lm.Append(rec)
	   if err != nil {
		   log.Fatal("occour error")
	   }
	   log.Printf("%d ", lsn)
   }
}

func createLogRecord(s string, n int, lm *LogMgr) []byte {
	npos := lm.logpage.MaxLength(len(s))
	b := make([]byte, npos + 4)
	p := file.NewPageFromByte(b)
	p.SetString(0, s)
	p.SetInt(npos, n)
	return b
}
