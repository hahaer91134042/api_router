package Impl

type IToString interface {
	ToString() string
}

/*
用來比較物件是否相同 目前只有在MAP上有用到
*/
type IEqual[T any] interface {
	Equal(v T) bool
}
