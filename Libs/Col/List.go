package Col

import (
	"Libs/Json"
)

type List[T any] []T

func (l List[T]) Length() int {
	return len(l)
}

func (l List[T]) First() T {
	return l.Get(0)
}

func (l List[T]) Last() T {
	return l.Get(l.Length() - 1)
}

func (l List[T]) Get(index int) (t T) {
	if index < 0 || index > (len(l)-1) {
		return
	}

	t = l[index]
	return
}

func (l List[T]) Add(data T) List[T] {
	// ret = l
	// if reflect.TypeFor[T]().Name() == reflect.TypeOf(data).Name() {
	// 	l.raw = append(l.raw, data)
	// }
	l = append(l, data)
	return l
}

func (l List[T]) RemoveAt(index int) List[T] {
	//l.raw = append(l.raw[:index], l.raw[index+1:]...)
	l = RemoveAt(l, index)
	return l
}

func (l List[T]) RemoveBy(fun func(T) bool) List[T] {
	var r_i = []T{}
	for _, v := range l {
		if !fun(v) {
			r_i = append(r_i, v)
		}
	}
	l = r_i
	return l
}

func (l List[T]) Where(fun func(T) bool) List[T] {

	return Where(l, fun)
}

// func (l *List[T]) Where(fun func(T) bool) []T {
// 	var newRaw = []T{}

// 	l.ForEach(func(i int, t T) {
// 		if fun(t) {
// 			newRaw = append(newRaw, t)
// 		}
// 	})

// 	return newRaw
// }

func (l List[T]) ForEach(fun func(int, T)) {
	ForEach(l, fun)
	// for i, v := range l.raw {
	// 	fun(i, v)
	// }
	//return l
}

func (l List[T]) ToString() string {
	return Json.String(l)
}

func (l List[T]) GetRaw() []T {
	return l
}

func NewList[T any](args ...T) List[T] {
	return args
}
