package gno

import (
	"reflect"
	"testing"
)

func TestCriteria(t *testing.T) {
	sql, args := Criteria{
		"name": "python",
	}.ToSql()
	expected := "name = ?"
	if sql != expected {
		t.Errorf("Criteria failed, got: %s, expected: %s.", sql, expected)
	}

	expectedArgs := []interface{}{"python"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Criteria failed, got: %v, expected: %v.", args, expectedArgs)
	}
}
