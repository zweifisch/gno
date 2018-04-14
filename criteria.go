package gno

import (
	"reflect"
	"strings"
)

type Criteria map[string]interface{}

type Expression interface {
	ToSql() (string, []interface{})
}

type And struct {
	Expressions []*Expression
}

func (criteria Criteria) ToSql() (string, []interface{}) {
	sql := []string{}
	args := []interface{}{}
	for field, value := range criteria {
		expr := ""
		if isList(value) {
			expr = field + " IN (" + strings.Join(Fill(len(value.([]interface{})), "?"), ",") + ")"
			args = append(args, value.([]interface{})...)
		} else {
			expr = field + " = ?"
			args = append(args, value)
		}
		sql = append(sql, expr)
	}
	return strings.Join(sql, " AND "), args
}

func isList(object interface{}) bool {
	val := reflect.ValueOf(object)
	return val.Kind() == reflect.Array || val.Kind() == reflect.Slice
}
