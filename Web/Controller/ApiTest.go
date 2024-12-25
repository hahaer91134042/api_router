package Controller

import (
	"HillWeb/Error"
	VM "HillWeb/ViewModel"
	"Libs/Log"
	"net/http"
)

type ApiTest struct {
	auth string
}

func (a *ApiTest) Path() string {
	return "/testapi"
}
func (a *ApiTest) Method() []string {
	return []string{http.MethodPost}
}

func (a *ApiTest) ReqHeaderAuth(req *http.Request, auth string) {
	Log.Warn.Print("ApiTest Get Auth=", auth)
	a.auth = auth
}

func (a *ApiTest) ExecuteAction(rw http.ResponseWriter, req *http.Request) VM.JSON {
	// data := struct {
	// 	Auth string `json:"auth"`
	// }{}
	Log.Warn.Print("ApiTest Run auth=", a.auth)
	return VM.ResData{
		Code: int(Error.Code.Success),
		Msg:  "This is ApiTest!!",
		Data: struct {
			Auth string `json:"auth"`
		}{
			Auth: a.auth,
		},
	}
}

func (a *ApiTest) OnClose() {

}
