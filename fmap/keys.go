package fmap

import (
	"reflect"
)

// Keys return an slice of all the keys of a map.
// Got slice that is not sorted.
func Keys(m interface{}) interface{} {
	if m == nil {
		return m
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return m
	}

	keys := reflect.MakeSlice(reflect.SliceOf(val.Type().Key()), 0, val.Len())
	if val.Len() == 0 {
		return keys.Interface()
	}

	iter := val.MapRange()
	for iter.Next() {
		keys = reflect.Append(keys, iter.Key())
	}

	return keys.Interface()
}
