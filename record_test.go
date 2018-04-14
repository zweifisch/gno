package gno

import (
	"reflect"
	"testing"
)

func TestUnzip(t *testing.T) {
	record := Record{
		"key":    "value",
		"number": 2,
	}
	keys, _ := record.Unzip()
	expected1 := []string{"key", "number"}
	expected2 := []string{"number", "key"}
	if !reflect.DeepEqual(keys, expected1) && !reflect.DeepEqual(keys, expected2) {
		t.Errorf("Unzip failed, got: %v, expected: %v.", keys, expected1)
	}
}
