package record

import (
	"testing"
)

func TestSchema(t *testing.T) {
	sch := NewSchema()
	t.Logf("initialize schema: %v", sch)

	// Test AddField method
	sch.AddField("title", 0, 10)
	if sch.Fields[0] != "title" {
		t.Errorf("AddField: expected field name: 'title', but got '%s'", sch.Fields[0])
	}
	if sch.Info["title"].Type != 0 || sch.Info["title"].Length != 10 {
		t.Errorf("AddField: expected type: 0 and length: 10, but got type: %d, length: %d", sch.Info["title"].Type, sch.Info["title"].Length)
	}

	// Test AddIntField method
	sch.AddIntField("age")
	if sch.Info["age"].Type != Integer {
		t.Errorf("AddIntField: expected type: 'Integer' but got %d", sch.Info["age"].Type)
	}

	// Test AddStringField method
	sch.AddStringField("name", 10)
	if sch.Info["name"].Type != Varchar || sch.Info["name"].Length != 10 {
		t.Errorf("AddStringField: expected type: 'Varchar' and length: 10, but got type: %d, length: %d", sch.Info["name"].Type, sch.Info["name"].Length)
	}

	//Test Add, AddAll methods
	newSchema := NewSchema()
	newSchema.AddIntField("newField1")
	newSchema.AddStringField("newField2", 20)
	sch.AddAll(*newSchema)
	if sch.Info["newField1"].Type != Integer || sch.Info["newField2"].Type != Varchar || sch.Info["newField2"].Length != 20 {
		t.Errorf("AddAll, Add: failed to add fields from another schema")
	}

	// Test HasField method
	if !sch.HasField("age") {
		t.Error("HasField: field 'age' should exist")
	}
	if sch.HasField("nonexistent") {
		t.Error("HasField: field 'nonexistent' should not exist")
	}

	// Test Type method
	fieldType, err := sch.Type("age")
	if err != nil {
		t.Errorf("Type: %v", err)
	}
	if fieldType != Integer {
		t.Errorf("Type: expected 'Integer' but got %d", fieldType)
	}

	// Test Length method
	fieldLength, err := sch.Length("name")
	if err != nil {
		t.Errorf("Length: %v", err)
	}
	if fieldLength != 10 {
		t.Errorf("Length: expected 10 but got %d", fieldLength)
	}
}
