package File

import "io/fs"

type Info struct {
	raw fs.DirEntry
	/*
		完整檔名
	*/
	FullName string
	/*
		主檔名
	*/
	Name string

	/*
		副檔名
	*/
	Sub string
}