package slice

import (
	"errors"
	"reflect"

	"github.com/barkimedes/go-deepcopy"

	"github.com/nastvood/cont/fslice/check"
)

// Make ...
func Make(n int, el interface{}) (interface{}, error) {
	if n < 0 {
		return nil, errors.New("n is negative")
	}

	if el == nil {
		return nil, errors.New("el is nil")
	}

	elType := reflect.TypeOf(el)
	newS := reflect.MakeSlice(reflect.SliceOf(elType), 0, n)
	if elType.Kind() == reflect.Ptr && reflect.ValueOf(el).IsNil() {
		for i := 0; i < n; i++ {
			newS = reflect.Append(newS, reflect.ValueOf(el))
		}

		return newS.Interface(), nil
	}

	for i := 0; i < n; i++ {
		elCopy, err := deepcopy.Anything(el)
		if err != nil {
			return nil, err
		}
		newS = reflect.Append(newS, reflect.ValueOf(elCopy))
	}

	return newS.Interface(), nil
}

// Init ...
func Init(n int, fn interface{}) (interface{}, error) {
	if n < 0 {
		return nil, errors.New("negative n")
	}

	fnType := reflect.TypeOf(fn)
	err := check.FuncSliceInit(reflect.TypeOf(fn))
	if err != nil {
		return nil, err
	}

	elType := fnType.Out(0)
	newS := reflect.MakeSlice(reflect.SliceOf(elType), 0, n)
	for i := 0; i < n; i++ {
		val := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(i)})[0]
		newS = reflect.Append(newS, val)
	}

	return newS.Interface(), nil
}
