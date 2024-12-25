package Controller

import (
	"App"
	"HillWeb/Error"
	"HillWeb/Model/Tpi_1"
	VM "HillWeb/ViewModel"
	"Libs/Log"
	Sql "SQL"
	"net/http"
)

type Test2 struct{}

func (t *Test2) Path() string {
	return "/api/test2"
}
func (t *Test2) Method() []string {
	return []string{http.MethodGet, http.MethodPost}
}

func (t *Test2) ExecuteAction(rw http.ResponseWriter, req *http.Request) VM.JSON {

	Log.Warn.Print("Test2 Api Db_Read->", App.Db_Read)

	//App.Db_Read.From(Tpi_1.User{})

	var cmd = Sql.From[Tpi_1.User](App.Db_Read)
		//Where(&Tpi_1.User{UserId: 1}).
		//Where(Tpi_1.User{UserId: 1})
	var count = cmd.Count()
	var data = cmd.First()

	Log.Warn.Print("Tpi_1.Use columns=>", data)
	Log.Warn.Print("Tpi_1.Use count=>", count)

	return VM.ResData{
		Code: int(Error.Code.InvalidAuth),
		Msg:  "haha",
		Data: data,
	}
}
func (t *Test2) OnClose() {

}
