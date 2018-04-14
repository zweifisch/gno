package gno

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	sql, args := Find("langs", &Criteria{
		"name": "python",
	}).ToSql()
	expected := "SELECT * FROM langs WHERE name = ?;"
	if sql != expected {
		t.Errorf("Criteria failed, got: %s, expected: %s.", sql, expected)
	}

	expectedArgs := []interface{}{"python"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Criteria failed, got: %v, expected: %v.", args, expectedArgs)
	}
}
