package fslice

import (
	"reflect"

	"github.com/nastvood/cont/fslice/check"
)

// Fold - func({s0, s1, ..., sn}, a, fn) is fn(sn, fn(..., fn(s1, fn(s0, a)))).
func Fold(s, zero, fn interface{}) interface{} {
	if s == nil {
		return zero
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return zero
	}

	err := check.FuncFold(reflect.TypeOf(fn), reflect.TypeOf(s).Elem())
	if err != nil {
		return zero
	}

	if val.Len() == 0 {
		return zero
	}

	res := reflect.ValueOf(zero)
	for i := 0; i < val.Len(); i++ {
		args := []reflect.Value{
			val.Index(i),
			res,
		}

		res = reflect.ValueOf(fn).Call(args)[0]
	}

	return res.Interface()
}
