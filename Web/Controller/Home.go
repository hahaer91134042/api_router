package Controller

import (
	"App"
	VM "HillWeb/ViewModel"
	"Libs/DateTime"
	"Libs/Log"
	"fmt"
	"net/http"
	"time"
)

type Home struct {
}

func (c *Home) Path() string {
	return "/"
}
func (c *Home) Method() []string {
	return []string{http.MethodGet}
}

func (c *Home) ExecuteAction(rw http.ResponseWriter, req *http.Request) VM.HTML {
	msg := "Hello, this is New Hill Web Api 123"
	//code := Error.Code.Success
	rspData := map[string]interface{}{
		"msg": msg,
	}
	Log.Debug.Print("Home Api run")

	// 版本
	rspData["version"] = App.Ver

	htmlStr := "<h1>Hellow This is Hill_Api</h1><br/>"
	htmlStr += fmt.Sprint("<span style=\"color:red;font-size:30px\">Version:", App.Ver, "</span><br/>")
	htmlStr += fmt.Sprint("<span style=\"font-size:15px\">Now Time:", time.Now().Format(DateTime.Format), "</span><br/>")

	return VM.HTML(htmlStr)

	// return struct {
	// 	// 錯誤代碼
	// 	Code int    `json:"code"`
	// 	Msg  string `json:"msg"`
	// 	// 資料
	// 	Data interface{} `json:"data"`
	// }{
	// 	Code: int(code),
	// 	Msg:  Enum.ToString(code),
	// 	Data: rspData,
	// }

	// rspBytes, _ := json.Marshal(struct {
	// 	// 錯誤代碼
	// 	Code int    `json:"code"`
	// 	Msg  string `json:"msg"`
	// 	// 資料
	// 	Data interface{} `json:"data"`
	// }{
	// 	Code: int(code),
	// 	Msg:  Enum.ToString(code),
	// 	Data: rspData,
	// })

}
func (h *Home) OnClose() {
	Log.Error.Print("Home On Close!")
}

// func Homepage(w http.ResponseWriter, r *http.Request) {

// 	msg := "Hello, this is manager_api_sv"
// 	code := Error.Code.Success
// 	rspData := map[string]interface{}{
// 		"msg": msg,
// 	}
// 	// 版本
// 	rspData["version"] = App.Ver

// 	rspBytes, _ := json.Marshal(struct {
// 		// 錯誤代碼
// 		Code int    `json:"code"`
// 		Msg  string `json:"msg"`
// 		// 資料
// 		Data interface{} `json:"data"`
// 	}{
// 		Code: int(code),
// 		Msg:  Enum.ToString(code),
// 		Data: rspData,
// 	})

// 	w.Write(rspBytes)
// }
