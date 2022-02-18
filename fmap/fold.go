package fmap

import (
	"reflect"

	"github.com/nastvood/cont/fmap/check"
)

// Fold - func(m, a, fn) is fn(kn, vn, fn(..., fn(k1, v2, fn(k0, v0, a)))).
func Fold(m, zero, fn interface{}) interface{} {
	if m == nil {
		return zero
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return zero
	}

	mType := reflect.TypeOf(m)
	err := check.FuncFold(reflect.TypeOf(fn), mType.Key(), mType.Elem(), reflect.TypeOf(zero))
	if err != nil {
		return zero
	}

	res := reflect.ValueOf(zero)
	iter := val.MapRange()
	for iter.Next() {
		args := []reflect.Value{
			iter.Key(),
			iter.Value(),
			res,
		}

		res = reflect.ValueOf(fn).Call(args)[0]
	}

	return res.Interface()
}
