package Json

import (
	"encoding/json"
)

func String(a any) string {
	s, e := json.Marshal(a)
	if e == nil {
		return string(s)
	}
	return ""
}

func Parse[T any](s string) *T {
	var a = new(T)
	e := json.Unmarshal([]byte(s), a)
	if e == nil {
		return a
	}
	return nil
}
