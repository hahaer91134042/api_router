package Web

import (
	VM "HillWeb/ViewModel"
	"net/http"
)

type IWebAction_Gorilla interface {
	Path() string
	Method() []string
	OnClose()
}

type IWebAuth_Gorilla interface {
	ReqHeaderAuth(req *http.Request, auth string)
}

type IWebAction_Json_Gorilla interface {
	ExecuteAction(rw http.ResponseWriter, req *http.Request) VM.JSON
}

type IWebAction_Html_Gorilla interface {
	ExecuteAction(rw http.ResponseWriter, req *http.Request) VM.HTML
}
