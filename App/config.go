package App

import (
	"Libs/Ext"
	Sql "SQL"
	"encoding/json"
)

type Mode int
type ModeLog Mode
type ModeHttp Mode

type Setting struct {
	Host struct { // 本地端設定
		Port    string `json:"port"`
		Version string `json:"version"`
		Mode    struct {
			Log  Mode `json:"log"`
			Http Mode `json:"http"`
		} `json:"mode"`
		Auth struct {
			CertFileName string `json:"certFileName"`
			KeyFileName  string `json:"keyFileName"`
		}
	} `json:"host"`
	Db struct { // Db設定
		Backstage struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			User     string `json:"user"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"backstage"`
	} `json:"db"`
	ExternalSites struct {
		DA struct {
			Url string `json:"url"`
		} `json:"DA"`
	} `json:"externalSites"`
}

func (s *Setting) ToSqlConfig() Sql.Config {
	return Sql.Config{
		User: s.Db.Backstage.User,
		Pwd:  s.Db.Backstage.Password,
		Host: s.Db.Backstage.Host,
		Port: s.Db.Backstage.Port,
		Db:   s.Db.Backstage.Database,
	}
}

const (
	Mode_Dev     ModeLog  = 1
	Mode_Release ModeLog  = 2
	Mode_Http    ModeHttp = 1
	Mode_Https   ModeHttp = 2
)

var Config = new(Setting)

var Ver = ""

var Web = struct {
	Port     string
	AuthCert string
	AuthKey  string
}{
	Port:     "",
	AuthCert: "",
	AuthKey:  "",
}

func InitMode(sb []byte) (err error) {
	err = json.Unmarshal(sb, Config)
	//Log.Debug.Printf("App InitMode Config=%+v", Config)
	//Log.Debug.Print("Config err=", err, " Config ver=", Config.Host.Version)

	if err == nil {
		Ver = Config.Host.Version
		Web.Port = Config.Host.Port
		Web.AuthCert = Config.Host.Auth.CertFileName
		Web.AuthKey = Config.Host.Auth.KeyFileName
	}
	//Log.Warn.Print("App.Ver ptr=", &Ver, " val=", Ver)
	// if Config.Host.Mode.Log == CML_Dev {
	// 	ModeLog = CML_Dev
	// }

	// if Config.Host.Mode.Http == CMH_Http {
	// 	ModeHttp = CMH_Http
	// }
	return
}

func GetMode[T ModeLog | ModeHttp]() T {
	b := Ext.TypeEqual[ModeLog](new(T))
	//Log.W(fmt.Sprintf("Mode equal log=%v", b))
	if b {
		return T(Config.Host.Mode.Log)
	}
	return T(Config.Host.Mode.Http)
}

func RunMode[T ModeLog | ModeHttp](in T, fn func()) {
	if m := GetMode[T](); m == in {
		fn()
	}
}
