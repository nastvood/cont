package slice

import (
	"reflect"

	"github.com/nastvood/cont/slice/util"
)

// ForAll ...
func ForAll(s, fn interface{}) bool {
	if s == nil {
		return true
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return false
	}

	elType := reflect.TypeOf(s).Elem()
	err := util.CheckFuncUnaryPredicate(reflect.TypeOf(fn), elType)
	if err != nil {
		return false
	}

	if val.Len() == 0 {
		return true
	}

	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
		if !v.Bool() {
			return false
		}
	}

	return true
}

// Exists ...
func Exists(s, fn interface{}) bool {
	if s == nil {
		return false
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return false
	}

	elType := reflect.TypeOf(s).Elem()
	err := util.CheckFuncUnaryPredicate(reflect.TypeOf(fn), elType)
	if err != nil {
		return false
	}

	if val.Len() == 0 {
		return false
	}

	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
		if v.Bool() {
			return true
		}
	}

	return false
}
