package Col

import "slices"

type ICollector[T any] interface {
	Get(i int) T
	ForEach(func(int, T))
	GetRaw() []T
}

type IQuery[T any, C ICollector[T]] interface {
	Where(func(T) bool) C
}

func Select[V any, T ICollector[V], R any](l T, fun func(V) R) []R {
	var ret = []R{}
	for _, v := range l.GetRaw() {
		ret = append(ret, fun(v))
	}
	return ret
}
func ForEach[V any, T ICollector[V]](l T, fun func(int, V)) {
	for i, v := range l.GetRaw() {
		fun(i, v)
	}
}

func Where[V any, T ICollector[V]](l T, fun func(V) bool) []V {
	var newRaw = []V{}
	for _, v := range l.GetRaw() {
		if fun(v) {
			newRaw = append(newRaw, v)
		}
	}
	return newRaw
}

func RemoveAt[T any, Q ICollector[T]](q Q, i int) []T {
	var raw = slices.Clone(q.GetRaw())
	//var newRaw = append(raw[:i], raw[i+1:]...)
	return append(raw[:i], raw[i+1:]...)
}
