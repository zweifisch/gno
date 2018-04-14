package gno

const DESC = "DESC"
const ASC = "ASC"

type FindStmt struct {
	table    string
	criteria Expression
	orderBy  string
	order    string
	offset   int
	limit    int
}

func Find(table string, expr Expression) *FindStmt {
	return &FindStmt{table: table, criteria: expr}
}

func (stmt *FindStmt) Limit(offset, limit int) *FindStmt {
	stmt.offset = offset
	stmt.limit = limit
	return stmt
}

func (stmt *FindStmt) OrderBy(field string, order string) *FindStmt {
	stmt.orderBy = field
	stmt.order = order
	return stmt
}

func (stmt *FindStmt) ToSql() (string, []interface{}) {
	var args []interface{}
	criteria, _args := stmt.criteria.ToSql()
	args = append(args, _args...)
	return "SELECT * FROM " + stmt.table + " WHERE " + criteria + ";", args
}
