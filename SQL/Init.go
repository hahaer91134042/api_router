package Sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg Config) (c *Client_Gorm, e error) {
	var cnnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.Db)

	//避免空
	if GormCfg == nil {
		GormCfg = &gorm.Config{}
	}

	sql, err := gorm.Open(mysql.Open(cnnStr), GormCfg)
	if err != nil {
		e = err
	}
	c = &Client_Gorm{gdb: sql}

	return
}
