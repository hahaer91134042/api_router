package File

import (
	"Libs/Log"
	"os"
	"strings"
)

func FromDir(path string) []Info {
	files, e := os.ReadDir(path)
	if e == nil {
		var l = []Info{}
		for _, f := range files {
			name := strings.Split(f.Name(), ".")
			l = append(l, Info{
				raw:      f,
				FullName: f.Name(),
				Name:     name[0],
				Sub:      name[1],
			})
		}
		return l
	}
	Log.Error.Print("File path:", path, " load error!")
	return nil
}
