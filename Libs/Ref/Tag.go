package Ref

import (
	"Libs/Col"
	"reflect"
)

func TagOf(a any, tag string) Col.Map[string, string] {
	m := Col.NewMap[string, string]()
	t := reflect.TypeOf(a)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get(tag)
		//fmt.Printf("Field: %s, Tag: %s\n", field.Name, tagValue)
		m.Add(field.Name, tagValue)
	}

	return m
}

func TagFor[T any](tag string) Col.Map[string, string] {
	var a T
	return TagOf(a, tag)
}
