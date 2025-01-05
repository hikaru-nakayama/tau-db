package record

type Layout struct {
	schema   *Schema
	offsets  map[string]int
	slotsize int
}
