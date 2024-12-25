package Ext

import "errors"

/*
arr=slice array
i=index
可能沒用處了
*/
func TryGet[T any](arr []T, i int) (T, error) {
	if len(arr) > i {
		return arr[i], nil
	}
	return *new(T), errors.New("out of index")

}
