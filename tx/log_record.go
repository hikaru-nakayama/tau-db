package tx

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
