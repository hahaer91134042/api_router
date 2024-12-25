package Ext

import (
	"reflect"
	"strings"
)

/*
check type equal
*/
func TypeEqual[A any](b any) bool {
	//去除ptr符號 只要type名稱
	as := strings.Replace(reflect.TypeFor[A]().String(), "*", "", -1)
	bs := strings.Replace(reflect.TypeOf(b).String(), "*", "", -1)

	//fmt.Printf("Ext A Type=%v B Type=%v\n", as, bs)
	return as == bs
}
