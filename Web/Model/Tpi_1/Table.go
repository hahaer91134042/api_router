package Tpi_1

type User struct {
	UserId  int64   `gorm:"column:userid"`
	Account string  `gorm:"column:account"`
	Money   float64 `gorm:"column:money"`
	Demo    string  `gorm:"column:demo"`
}

func (u User) SqlTable() (string, []any) {
	return "user", []any{}
}

func (u User) SelectColumn() []string {
	return []string{"userid", "account", "money", "demo"}
}
