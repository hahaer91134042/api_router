package Col

import (
	"Libs/Impl"
	"Libs/Json"
	"reflect"
)

type Key_Value[K any, V any] struct {
	Key   K
	Value V
}

type Map[K any, V any] []Key_Value[K, V]

func (m Map[K, V]) Length() int {
	return len(m)
}

func (m Map[K, V]) First() Key_Value[K, V] {
	return m.Get(0)
}

func (m Map[K, V]) Last() Key_Value[K, V] {
	return m.Get(m.Length() - 1)
}

func (m Map[K, V]) GetValue(index int) (v V) {
	if index < 0 || index > (len(m)-1) {
		return
	}
	v = m[index].Value
	return
}

func (m Map[K, V]) Get(index int) (p Key_Value[K, V]) {
	if index < 0 || index > (len(m)-1) {
		return
	}
	p = m[index]
	return
}

func (m Map[K, V]) Find(fun func(K) bool) (v V) {
	for _, pair := range m {
		if fun(pair.Key) {
			v = pair.Value
			return
		}
	}
	//e = errors.New("Pair Not Found")
	return
}

func (m Map[K, V]) FindPair(fun func(K) bool) (p Key_Value[K, V]) {
	for _, pair := range m {
		if fun(pair.Key) {
			p = pair
			return
		}
	}
	//e = errors.New("Pair Not Found")
	return
}

func (m Map[K, V]) Add(k K, v V) Map[K, V] {
	m = append(m, Key_Value[K, V]{Key: k, Value: v})
	return m
}

func (m Map[K, V]) Where(fun func(Key_Value[K, V]) bool) Map[K, V] {

	return Where(m, fun)
}

// func (l *Map[K, V]) Where(fun func(Key_Value[K, V]) bool) []Key_Value[K, V] {
// 	var newRaw = []Key_Value[K, V]{}

// 	l.ForEach(func(i int, t Key_Value[K, V]) {
// 		if fun(t) {
// 			newRaw = append(newRaw, t)
// 		}
// 	})

// 	return newRaw
// }

func (m Map[K, V]) ForEach(fun func(int, Key_Value[K, V])) {
	ForEach(m, fun)
	// for i, v := range l.raw {
	// 	fun(i, v)
	// }
	// return l
}

//	func (m *Map[K, V]) RemoveAt(i int) *Map[K, V] {
//		m.raw = append(m.raw[:i], m.raw[i+1:]...)
//		return m
//	}

/*
Remove First element function
*/
func (m Map[K, V]) Remove(k K) Map[K, V] {
	kVal := reflect.ValueOf(k)
	var r_i = []Key_Value[K, V]{}
	if kVal.Type().Implements(reflect.TypeFor[Impl.IEqual[K]]()) {
		eq := reflect.ValueOf(k).MethodByName("Equal")
		for _, pair := range m {
			ret := eq.Call([]reflect.Value{reflect.ValueOf(pair.Key)})
			if !ret[0].Bool() {
				r_i = append(r_i, pair)
			}
		}
	} else {
		//var fun func(reflect.Value) bool
		for _, pair := range m {

			switch kVal.Kind() {
			case reflect.String:
				if kVal.String() != reflect.ValueOf(pair.Key).String() {
					r_i = append(r_i, pair)
				}

			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				if kVal.Int() != reflect.ValueOf(pair.Key).Int() {
					r_i = append(r_i, pair)
				}
			case reflect.Float32, reflect.Float64:
				if kVal.Float() != reflect.ValueOf(pair.Key).Float() {
					r_i = append(r_i, pair)
				}
			}

			// if r_i > -1 {
			// 	break
			// }

		}
		// if fun(reflect.ValueOf(k)) {
		// 	r_i = i
		// }
	}

	m = r_i

	// if r_i > -1 {
	// 	m.raw = RemoveAt(m, r_i)
	// }
	return m
}

func (m Map[K, V]) RemoveBy(fun func(K) bool) Map[K, V] {
	var r_i = []Key_Value[K, V]{}
	for _, pair := range m {
		if fun(pair.Key) {
			r_i = append(r_i, pair)
		}
	}

	m = r_i

	return m
}

func (m Map[K, V]) ToString() string {
	return Json.String(m)
}

func (l Map[K, V]) GetRaw() []Key_Value[K, V] {
	return l
}

func NewMap[K any, V any]() Map[K, V] {
	return []Key_Value[K, V]{}
}

func NewMapRaw[K int | int16 | int32 | int64 | float32 | float64 | string, V any](raw map[K]V) Map[K, V] {
	var l = NewList[Key_Value[K, V]]()
	for k, v := range raw {
		l.Add(Key_Value[K, V]{Key: k, Value: v})
	}
	return l.GetRaw()
}
func NewMapFrom[K any, V any](args ...[]any) Map[K, V] {
	var l = NewList[Key_Value[K, V]]()
	for _, a := range args {
		l.Add(Key_Value[K, V]{
			Key:   a[0].(K),
			Value: a[1].(V),
		})
	}
	return l.GetRaw()
}
