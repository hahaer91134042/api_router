package Web

import (
	"HillWeb/Controller"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var ctrl = []IWebAction_Gorilla{
	new(Controller.Home),
	new(Controller.ApiTest),
	new(Controller.Test2),
}

func InitCors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	return cors
}

func InitFromGorilla() (r *mux.Router) {
	r = mux.NewRouter()
	for _, c := range ctrl {
		r.HandleFunc(
			c.Path(),
			func(w http.ResponseWriter, r *http.Request) {
				cType := reflect.TypeOf(c)

				if cType.Implements(reflect.TypeFor[IWebAuth_Gorilla]()) {
					c.(IWebAuth_Gorilla).ReqHeaderAuth(r, r.Header.Get("Authorization"))
				}

				// c.ExecuteAction(w, r)

				if cType.Implements(reflect.TypeFor[IWebAction_Json_Gorilla]()) {
					rspByte, err := json.Marshal(c.(IWebAction_Json_Gorilla).ExecuteAction(w, r))
					if err == nil {
						w.Header().Set("Content-Type", "application/json; charset=utf-8")
						w.Write(rspByte)
					}
				} else if cType.Implements(reflect.TypeFor[IWebAction_Html_Gorilla]()) {
					htmlStr := c.(IWebAction_Html_Gorilla).ExecuteAction(w, r)

					//h := []byte(html.UnescapeString(string(htmlStr)))
					//h := []byte(html.EscapeString(string(htmlStr)))
					h := []byte(string(htmlStr))

					w.Header().Set("Content-Type", "text/html; charset=utf-8")
					w.Write(h)
				}

				// if bs != nil {
				// 	w.Write(bs)
				// }

				defer c.OnClose()

			}).Methods(c.Method()...)
	}

	return
}
