package Sql

import (
	"Libs/Col"
	"Libs/Log"
	"Libs/Ref"
	SqlTag "SQL/Tag"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type query_arg struct {
	query any
	args  []any
}

type OrderBy struct {
	Name string
	Desc bool
}

func From[RET ITable](client *Client_Gorm) Cmd_Gorm[RET] {
	return Cmd_Gorm[RET]{Client: client}
}

type Cmd_Gorm[TAB ITable] struct {
	Client *Client_Gorm
	where  *query_arg
	join   *query_arg
	group  *string
	order  []OrderBy
	limit  int
}

func (c Cmd_Gorm[TAB]) Where(q any, arg ...any) Cmd_Gorm[TAB] {
	c.where = &query_arg{query: q, args: arg}
	return c
}

func (c Cmd_Gorm[TAB]) Join(q string, arg ...any) Cmd_Gorm[TAB] {
	c.join = &query_arg{query: q, args: arg}
	return c
}

func (c Cmd_Gorm[TAB]) GroupBy(group string) Cmd_Gorm[TAB] {
	c.group = &group
	return c
}
func (c Cmd_Gorm[TAB]) OrderBy(order ...OrderBy) Cmd_Gorm[TAB] {
	c.order = append(c.order, order...)
	return c
}

func (c Cmd_Gorm[TAB]) Limit(l int) Cmd_Gorm[TAB] {
	c.limit = l
	return c
}

func (c Cmd_Gorm[TAB]) Rows() Col.List[TAB] {
	var res = make([]TAB, 0)
	var db = c.queryDb()
	db.Scan(&res)

	return Col.NewList(res...)
}

func (c Cmd_Gorm[TAB]) RowsCount() (count int64, row Col.List[TAB]) {
	row = c.Rows()
	count = int64(row.Length())
	return
}

func (c Cmd_Gorm[TAB]) First() TAB {
	var res TAB

	var db = c.queryDb()
	db.First(&res)
	return res
}

func (c Cmd_Gorm[TAB]) Count() (count int64) {
	var db = c.queryDb()
	db.Count(&count)
	return
}

func (c Cmd_Gorm[TAB]) queryDb() *gorm.DB {
	var tab TAB
	tName, tArg := tab.SqlTable()
	//Log.Info.Print("Cmd Tab Name->", tName, " args->", tArg)
	var db = c.Client.gdb.Table(tName, tArg...).Select(c.getSelect())

	// Log.Info.Print("Cmd where->", c.where)
	// Log.Info.Print("Cmd join->", c.join)

	if c.where != nil {
		db = db.Where(c.where.query, c.where.args...)
	}
	if c.join != nil {
		db = db.Joins(c.join.query.(string), c.join.args...)
	}

	if c.group != nil {
		db = db.Group(*c.group)
	}

	if len(c.order) > 0 {
		var order = clause.OrderBy{}
		for _, o := range c.order {
			order.Columns = append(order.Columns, clause.OrderByColumn{
				Column: clause.Column{Name: o.Name},
				Desc:   o.Desc,
			})
		}
		db = db.Order(order)
	}

	if c.limit > 0 {
		db = db.Limit(c.limit)
	}
	return db
}

func (c Cmd_Gorm[TAB]) getSelect() []string {

	if reflect.TypeFor[TAB]().Implements(reflect.TypeFor[ISelectColumn]()) {
		return reflect.New(reflect.TypeFor[TAB]()).Interface().(ISelectColumn).SelectColumn()
	} else {
		var tab TAB
		var colList = Col.NewList[string]()
		SqlTag.ParseGorm(Ref.TagOf(tab, "gorm")).
			ForEach(func(i int, p Col.Key_Value[string, SqlTag.Gorm]) {
				Log.Debug.Print("Field:", p.Key, " gorm column:", p.Value.Column)
				colList.Add(p.Value.Column)
			})
		return colList.GetRaw()
	}
}
