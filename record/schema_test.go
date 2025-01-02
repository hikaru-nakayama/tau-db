package record

import (
	"testing"
)

func TestSchema(t *testing.T) {
	sch := NewSchema()
	t.Logf("initialize schema: %v", sch)

	sch.AddField("title", 0, 10)

	if sch.Fields[0] != "title" {
		t.Errorf("expected: title, but got %s", sch.Fields[0])
	}

	if sch.Info["title"].Type != 0 || sch.Info["title"].Length != 10 {
		t.Errorf("expected: type is 0 and length is 10, but got type: %d, length: %d", sch.Info["title"].Type, sch.Info["title"].Length)
	}

	sch.AddIntFiled("age")

	if sch.Info["age"].Type != Integer {
		t.Errorf("expected: type is Integer but got %d", sch.Info["name"].Type)
	}

	sch.AddStringFiled("name", 10)

	if sch.Info["name"].Type != Varchar || sch.Info["name"].Length != 10 {
		t.Errorf("expected: type is Varchar and length is 10, but got type: %d, length: %d", sch.Info["name"].Type, sch.Info["name"].Length)
	}
}
