package gno

import (
	"fmt"
	"strings"
)

type InsertStmt struct {
	table  string
	record *Record
	update *Record
}

func Insert(table string, record *Record) *InsertStmt {
	return &InsertStmt{table: table, record: record}
}

func (stmt *InsertStmt) OnDuplicateKeyUpdate(record *Record) *InsertStmt {
	stmt.update = record
	return stmt
}

func (stmt *InsertStmt) ToSql() (string, []interface{}) {
	keys, args := stmt.record.Unzip()
	valuePlaceholder := "(" + strings.Join(MapString(keys, func(x string) string {
		return "?"
	}), ",") + ")"
	sql := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES %s", stmt.table, strings.Join(keys, ","), valuePlaceholder)

	if stmt.update != nil {
		keys, values := stmt.update.Unzip()
		sql = sql + " ON DUPLICATE KEY UPDATE " + strings.Join(MapString(keys, func(column string) string {
			return column + "=?"
		}), ",")
		args = append(args, values...)
	}
	return sql + ";", args
}

func Upsert(tbl string, record *Record) *InsertStmt {
	return Insert(tbl, record).OnDuplicateKeyUpdate(record)
}
