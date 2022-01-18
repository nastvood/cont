package check

import (
	"fmt"
	"reflect"

	"github.com/nastvood/cont/errors"
)

func FuncMap(fnType, keyType, valueType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 2 {
		return errors.NewNumInError(2)
	}

	if fnType.NumOut() != 2 {
		return errors.NewNumOutError(2)
	}

	if fnType.In(0) != keyType {
		return errors.NewInTypeError(0, keyType, fnType.In(0))
	}

	if fnType.In(1) != valueType {
		return errors.NewInTypeError(0, valueType, fnType.In(1))
	}

	return nil
}

func FuncFold(fnType, keyType, valueType, zeroType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	fmt.Println(fnType.NumIn())
	if fnType.NumIn() != 3 {
		return errors.NewNumInError(3)
	}

	if fnType.NumOut() != 1 {
		return errors.NewNumOutError(1)
	}

	if fnType.In(0) != keyType {
		return errors.NewInTypeError(0, keyType, fnType.In(0))
	}

	if fnType.In(1) != valueType {
		return errors.NewInTypeError(0, valueType, fnType.In(1))
	}

	if fnType.Out(0) != fnType.In(2) {
		return errors.NewOutTypeError(0, fnType.In(2), fnType.Out(0))
	}

	if fnType.Out(0) != zeroType {
		return errors.NewOutTypeError(0, zeroType, fnType.Out(0))
	}

	return nil
}
