package Sql

import (
	"gorm.io/gorm"
)

type Client_Gorm struct {
	gdb *gorm.DB
}

func (c *Client_Gorm) From(tab string, arg ...any) fromTable_Gorm {

	return fromTable_Gorm{
		db: c.gdb.Table(tab, arg...),
	}
}

// func (c *Client_Gorm) From2(tab ITable) {
// 	tName, arg := tab.SqlTable()
// 	Log.Debug.Print("Sql Table Name->", tName)
// 	sql := c.gdb.Table(tName, arg...)

// 	if reflect.TypeOf(tab).Implements(reflect.TypeFor[ISelectColumn]()) {
// 		sql = sql.Select(tab.(ISelectColumn).SelectColumn())
// 	} else {
// 		var colList = Col.NewList[string]()
// 		SqlTag.ParseGorm(*Ref.TagOf(tab, "gorm")).
// 			ForEach(func(i int, p Col.Key_Value[string, SqlTag.Gorm]) {
// 				Log.Debug.Print("Field:", p.Key, " gorm column:", p.Value.Column)
// 				colList.Add(p.Value.Column)
// 			})
// 		sql = sql.Select(colList.GetRaw())
// 	}

// }

/*
暫時別用 開發中
models.DbRead.

	Table("game_bet_result").
	Select([]string{"OrderID", "LiveOrderID", "PlayerName", "PlayerID", "GameID", "game_name", "BetMoney", "SUM(game_bet_result.PrizeMoney) AS PrizeMoney", "Time", "LiveOrderID"}).
	Joins("left join `managedb`.`sync_game_info` ON sync_game_info.game_id = game_bet_result.GameID AND sync_game_info.language_id = 1").
	Where("PlayerName = ? AND GameID = ?", reqAccount, reqGameId).
	Group("OrderID").
	Order("Time DESC").
	Limit(5).
	Scan(&rslts)
*/
type fromTable_Gorm struct {
	db *gorm.DB
}

func (f fromTable_Gorm) Select(query any, arg ...any) fromTable_Gorm {
	f.db = f.db.Select(query, arg...)
	return f
}

func (f fromTable_Gorm) Where(query any, arg ...any) fromTable_Gorm {
	f.db = f.db.Where(query, arg...)
	return f
}

func (f fromTable_Gorm) Join(query string, arg ...any) fromTable_Gorm {
	f.db = f.db.Joins(query, arg...)
	return f
}

func (f fromTable_Gorm) GroupBy(query string) fromTable_Gorm {
	f.db = f.db.Group(query)
	return f
}

func (f fromTable_Gorm) OrderBy(query any) fromTable_Gorm {
	f.db = f.db.Order(query)
	return f
}

func (f fromTable_Gorm) Limit(val int) fromTable_Gorm {
	f.db = f.db.Limit(val)
	return f
}