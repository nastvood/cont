package fslice

import (
	"reflect"

	"github.com/nastvood/cont/fslice/check"
)

// Map - apply function fn to elements of slice s.
func Map(s, fn interface{}) interface{} {
	if s == nil {
		return s
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return s
	}

	fnType := reflect.TypeOf(fn)
	if err := check.FuncMap(fnType, reflect.TypeOf(s).Elem()); err != nil {
		return s
	}

	outType := fnType.Out(0)
	newS := reflect.MakeSlice(reflect.SliceOf(outType), 0, val.Cap())
	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
		newS = reflect.Append(newS, v)
	}

	return newS.Interface()
}

// Mapi ...
func Mapi(s, fn interface{}) interface{} {
	if s == nil {
		return s
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return s
	}

	fnType := reflect.TypeOf(fn)
	if err := check.FuncMapi(fnType, reflect.TypeOf(s).Elem()); err != nil {
		return s
	}

	outType := fnType.Out(0)
	newS := reflect.MakeSlice(reflect.SliceOf(outType), 0, val.Cap())
	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(i), val.Index(i)})[0]
		newS = reflect.Append(newS, v)
	}

	return newS.Interface()
}
