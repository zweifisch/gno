package gno

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	sql, data := Insert("tbl1", &Record{
		"name": "python",
	}).ToSql()
	expected := "INSERT INTO `tbl1` (name) VALUES (?);"
	if sql != expected {
		t.Errorf("Insert failed, got: %s, expected: %s.", sql, expected)
	}

	expectedData := []interface{}{"python"}
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Insert failed, got: %v, expected: %v.", data, expectedData)
	}
}

func TestUpsert(t *testing.T) {
	sql, data := Upsert("tbl1", &Record{
		"name": "python",
	}).ToSql()
	expected := "INSERT INTO `tbl1` (name) VALUES (?) ON DUPLICATE KEY UPDATE name=?;"
	if sql != expected {
		t.Errorf("Save failed, got: %s, expected: %s.", sql, expected)
	}
	expectedData := []interface{}{"python", "python"}
	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Save failed, got: %v, expected: %v.", data, expectedData)
	}
}
