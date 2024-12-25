package Sql

import "gorm.io/gorm"

/*
可以提供給外部修改使用
預設值是無設定的
*/
var GormCfg = &gorm.Config{}

type Config struct {
	Host string
	Port string
	Db   string
	User string
	Pwd  string
}
