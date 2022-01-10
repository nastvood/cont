package slice

import (
	"reflect"

	"github.com/nastvood/cont/slice/util"
)

// Fold ...
func Fold(s, zero, fn interface{}) interface{} {
	if s == nil {
		return zero
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return zero
	}

	err := util.CheckFuncFold(reflect.TypeOf(fn), reflect.TypeOf(s).Elem(), reflect.TypeOf(zero))
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
