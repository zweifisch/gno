package gno

type DeleteStmt struct {
	table    string
	criteria Expression
	offset   int
	limit    int
}

func Delete(table string, expr Expression) *DeleteStmt {
	return &DeleteStmt{table: table, criteria: expr}
}

func (stmt *DeleteStmt) Limit(offset, limit int) *DeleteStmt {
	stmt.offset = offset
	stmt.limit = limit
	return stmt
}

func (stmt *DeleteStmt) ToSql() (string, []interface{}) {
	var args []interface{}
	criteria, _args := stmt.criteria.ToSql()
	args = append(args, _args...)
	return "DELETE FROM " + stmt.table + " WHERE " + criteria + ";", args
}
