package check

import (
	"reflect"

	"github.com/nastvood/cont/errors"
	"github.com/nastvood/cont/internal/pkg/util"
)

func FuncFold(fnType, elemType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 2 {
		return errors.NewNumInError(2)
	}

	if fnType.NumOut() != 1 {
		return errors.NewNumOutError(1)
	}

	if fnType.In(0) != elemType {
		return errors.NewInTypeError(0, elemType, fnType.In(0))
	}

	if fnType.In(1) != fnType.Out(0) {
		return errors.NewInTypeError(1, fnType.In(1), fnType.Out(0))
	}

	// if fnType.Out(0) != foldType {
	// 	return errors.NewOutTypeError(0, foldType, fnType.Out(0))
	// }

	return nil
}

func FuncMapi(fnType, elemType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 2 {
		return errors.NewNumInError(2)
	}

	if fnType.NumOut() != 1 {
		return errors.NewNumOutError(1)
	}

	if fnType.In(0).Kind() != reflect.Int {
		return errors.NewInTypeError(0, util.IntType, fnType.In(0))
	}

	if fnType.In(1) != elemType {
		return errors.NewInTypeError(0, elemType, fnType.In(1))
	}

	return nil
}

func FuncFilterMap(fnType, elemType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 1 {
		return errors.NewNumInError(1)
	}

	if fnType.NumOut() != 2 {
		return errors.NewNumOutError(2)
	}

	if fnType.In(0) != elemType {
		return errors.NewInTypeError(0, elemType, fnType.In(0))
	}

	if fnType.Out(1).Kind() != reflect.Bool {
		return errors.NewOutTypeError(1, util.BoolType, fnType.Out(1))
	}

	return nil
}

func FuncSliceInit(fnType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 1 {
		return errors.NewNumInError(1)
	}

	if fnType.In(0).Kind() != reflect.Int {
		return errors.NewInTypeError(0, util.IntType, fnType.In(0))
	}

	if fnType.NumOut() != 1 {
		return errors.NewNumOutError(1)
	}

	return nil
}

func FuncUnaryPredicate(fnType, elemType reflect.Type) error {
	return funcOneOne(fnType, elemType, reflect.TypeOf(false))
}

func FuncMap(fnType, elemType reflect.Type) error {
	return funcOneOne(fnType, elemType, nil)
}

func funcOneOne(fnType, elemType, resType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.NewNotFuncError()
	}

	if fnType.NumIn() != 1 {
		return errors.NewNumInError(1)
	}

	if fnType.NumOut() != 1 {
		return errors.NewNumOutError(1)
	}

	if fnType.In(0) != elemType {
		return errors.NewInTypeError(0, elemType, fnType.In(0))
	}

	if resType != nil {
		if fnType.Out(0) != resType {
			return errors.NewOutTypeError(0, resType, fnType.Out(0))
		}
	}

	return nil
}
