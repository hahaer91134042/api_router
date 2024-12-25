package VM

type ResData struct {
	// 錯誤代碼
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// 資料
	Data interface{} `json:"data"`
}
