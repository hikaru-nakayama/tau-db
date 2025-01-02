package record

import (
	"errors"
)

type FiledInfo struct {
	Type   int
	Length int
}

type Schema struct {
	Fields []string
	Info   map[string]*FiledInfo
}

func NewSchema() *Schema {
	return &Schema{
		Fields: make([]string, 0),
		Info:   make(map[string]*FiledInfo),
	}
}

func (sch *Schema) AddField(field_name string, field_type int, length int) {
	sch.Fields = append(sch.Fields, field_name)
	sch.Info[field_name] = &FiledInfo{
		Type:   field_type,
		Length: length,
	}
}

func (sch *Schema) AddIntField(field_name string) {
	// do not use length when type is integer
	sch.AddField(field_name, Integer, 0)

}

func (sch *Schema) AddStringField(field_name string, length int) {
	sch.AddField(field_name, Varchar, length)
}

func (sch *Schema) Add(field_name string, schema Schema) {
	field_type := schema.Info[field_name].Type
	length := schema.Info[field_name].Length

	sch.AddField(field_name, field_type, length)
}

func (sch *Schema) AddAll(schema Schema) {
	for _, field_name := range schema.Fields {
		sch.Add(field_name, schema)
	}
}

func (sch *Schema) HasField(field_name string) bool {
	for _, field := range sch.Fields {
		if field == field_name {
			return true
		}
	}
	return false
}

func (sch *Schema) Type(field_name string) (int, error) {
	info, ok := sch.Info[field_name]
	if !ok {
		return 0, errors.New("filed not found")
	}

	return info.Type, nil
}

func (sch *Schema) Length(field_name string) (int, error) {
	info, ok := sch.Info[field_name]
	if !ok {
		return 0, errors.New("filed not found")
	}

	return info.Length, nil
}

const (
	Integer = 1
	Varchar = 2
)
