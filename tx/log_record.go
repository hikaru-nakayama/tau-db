package tx

type LogRecord interface {
	Op() int
	TxNumber() int
	Undo(tx +Transaction) 
}


