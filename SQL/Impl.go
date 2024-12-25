package Sql

type ITable interface {
	SqlTable() (string, []any)
}

type ISelectColumn interface {
	SelectColumn() []string
}
