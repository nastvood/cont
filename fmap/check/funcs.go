package check

import (
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
