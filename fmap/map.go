package fmap

import (
	"reflect"

	"github.com/nastvood/cont/fmap/check"
)

func Map(m, fn interface{}) interface{} {
	if m == nil {
		return m
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return m
	}

	fnType := reflect.TypeOf(fn)
	mType := reflect.TypeOf(m)
	if err := check.FuncMap(fnType, mType.Key(), mType.Elem()); err != nil {
		return m
	}

	outKeyType := fnType.Out(0)
	outValueType := fnType.Out(1)
	newMap := reflect.MakeMap(reflect.MapOf(outKeyType, outValueType))

	iter := val.MapRange()
	for iter.Next() {
		args := []reflect.Value{
			iter.Key(),
			iter.Value(),
		}

		vals := reflect.ValueOf(fn).Call(args)
		newMap.SetMapIndex(vals[0], vals[1])
	}

	return newMap.Interface()
}
