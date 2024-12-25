package App

import (
	"Libs/Log"
	Sql "SQL"
	"log"
)

var Db_Read *Sql.Client_Gorm
var Db_Write *Sql.Client_Gorm

func InitDb() {
	var err error
	Db_Read, err = Sql.Connect(Config.ToSqlConfig())
	Log.Debug.Print("Db_Read->", Db_Read)
	if err != nil {
		Log.Error.Print("App Db_Read link sql failed. ", err)
		log.Fatalln()
	}
	Db_Write, err = Sql.Connect(Config.ToSqlConfig())
	Log.Debug.Print("Db_Write->", Db_Write)
	if err != nil {
		Log.Error.Print("App Db_Write link sql failed. ", err)
		log.Fatalln()
	}
}
