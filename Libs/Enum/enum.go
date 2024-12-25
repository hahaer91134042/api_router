package Enum

import (
	"reflect"
)

//type Enum int

type IEnum interface {
	FlagMap() map[int]string
}

// type TestEnum int

// func (e TestEnum) FlagMap() map[int]string {
// 	return map[int]string{1: "Test", 2: "Test2"}
// }

// const (
// 	Test TestEnum = 1
// )

//	func ToString[T Enum]() string {
//		return ""
//	}
func ToString(a IEnum) string {
	//t := reflect.TypeOf(a)

	// if b := t.Implements(reflect.TypeFor[IEnum]()); !b {
	// 	return ""
	// }

	var m = a.FlagMap()
	// fmt.Println("Map:", m)
	// fmt.Println("Type:", t, "Value:", m[int(v.Int())])
	//output = m[int(reflect.ValueOf(a).Int())]
	return m[int(reflect.ValueOf(a).Int())]

}

// func Init() {

// 	fmt.Println("Test Enum V:", ToString(Test))
// 	//Test.ToString()
// }

// func (s Enum) ToString() string {
//     strings := [...]string{"PENDING", "ACTIVE"}

//     // prevent panicking in case of status is out-of-range
//     if s < pending || s > active {
//         return "Unknown"
//     }

//     return strings[s]
// }
