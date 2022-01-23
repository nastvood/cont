package fslice

import (
	"reflect"

	"github.com/nastvood/cont/fslice/check"
)

// Map - func({s0, s1, ..., sn}, fn), applies function fn to s0, s1, ..., sn
// and builds the slice {fn(a0), fn(a1), ..., fn(an)} with the results returned by fn.
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

// Mapi - func({s0, s1, ..., sn}, fn) will build the slice containing
// {fn(0, s0), fn(1, s1), ..., fn(n, sn)} where s0..sn are the elements of the slice s.
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
