package tx

import "github.com/hikaru-nakayama/tau-db.git/file"

type LogRecord interface {
	Op() int
	TxNumber() int
	Undo(tx *Transaction)
}

const (
	CHECKPOINT = iota
	START
	COMMIT
	ROLLBACK
	SETINT
	SETSTRING
)

func CreateLogRecord(bytes []byte) LogRecord {
	p := file.NewPageFromByte(bytes)
	switch p.GetInt(0) {
	case CHECKPOINT:
		return NewCheckPointRecord(p)
	case START:
		return NewStartRecord(p)
	case COMMIT:
		return NewCommitRecord(p)
	case ROLLBACK:
		return NewRollBackRecord(p)
	case SETINT:
		return NewSetIntRecord(p)
	case SETSTRING:
		return NewSetStringRecord(p)
	default:
		return nil
	}
}
