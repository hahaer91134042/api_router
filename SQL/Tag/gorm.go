package SqlTag

import (
	"Libs/Col"
	"strings"
)

type Gorm struct {
	Column  string
	Type    string
	Comment string
}

func ParseGorm(m Col.Map[string, string]) Col.Map[string, Gorm] {
	var gMap = Col.NewMap[string, Gorm]()
	m.Where(func(p Col.Key_Value[string, string]) bool {
		return p.Value != "-"
	}).ForEach(func(i int, p Col.Key_Value[string, string]) {
		var args = strings.Split(p.Value, ";")
		var g = Gorm{}
		for _, arg := range args {
			var gp = strings.Split(arg, ":")
			switch gp[0] {
			case "column":
				g.Column = gp[1]
			case "type":
				g.Type = gp[1]
			case "comment":
				g.Comment = gp[1]
			}
		}
		gMap.Add(p.Key, g)
	})

	return gMap
}
