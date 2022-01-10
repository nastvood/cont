package util

import (
	"errors"
	"fmt"
	"reflect"
)

func CheckFuncFold(fnType, elemType, resType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.New("fn is not func")
	}

	if fnType.NumIn() != 2 {
		return errors.New("number of args is not equal 1")
	}

	if fnType.NumOut() != 1 {
		return errors.New("number of outputs is not equal 1")
	}

	if fnType.In(0) != elemType {
		return fmt.Errorf("first arg has wrong type, actual %s, need %s", fnType.In(0), elemType)
	}

	if fnType.In(1) != resType {
		return fmt.Errorf("second arg has wrong type, actual %s, need %s", fnType.In(1), resType)
	}

	if resType != nil {
		if fnType.Out(0) != resType {
			return fmt.Errorf("output has wrong type, actual %s, need %s", fnType.Out(0), resType)
		}
	}

	return nil
}

func CheckFuncMapi(fnType, elemType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.New("fn is not func")
	}

	if fnType.NumIn() != 2 {
		return errors.New("number of args is not equal 2")
	}

	if fnType.NumOut() != 1 {
		return errors.New("number of outputs is not equal 1")
	}

	if fnType.In(0).Kind() != reflect.Int {
		return fmt.Errorf(fmt.Sprintf("arg(0) has wrong type, actual %s, need %s", fnType.In(1), elemType))
	}

	if fnType.In(1) != elemType {
		return fmt.Errorf(fmt.Sprintf("arg(1) has wrong type, actual %s, need %s", fnType.In(1), elemType))
	}

	return nil
}

func CheckFuncFilterMap(fnType, elemType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.New("fn is not func")
	}

	if fnType.NumIn() != 1 {
		return errors.New("number of args is not equal 1")
	}

	if fnType.NumOut() != 2 {
		return errors.New("number of outputs is not equal 2")
	}

	if fnType.In(0) != elemType {
		return fmt.Errorf(fmt.Sprintf("arg has wrong type, actual %s, need %s", fnType.In(0), elemType))
	}

	if fnType.Out(1).Kind() != reflect.Bool {
		return fmt.Errorf("output has wrong type, actual %s, need bool", fnType.Out(1))
	}

	return nil
}

func CheckFuncSliceInit(fnType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.New("fn is not func")
	}

	if fnType.NumIn() != 1 {
		return errors.New("number of args is not equal 2")
	}

	if fnType.In(0).Kind() != reflect.Int {
		return fmt.Errorf(fmt.Sprintf("arg(0) has wrong type, actual %s, need int", fnType.In(1)))
	}

	if fnType.NumOut() != 1 {
		return errors.New("number of outputs is not equal 1")
	}

	return nil
}

func CheckFuncUnaryPredicate(fnType, elemType reflect.Type) error {
	return checkFuncOneOne(fnType, elemType, reflect.TypeOf(false))
}

func CheckFuncMap(fnType, elemType reflect.Type) error {
	return checkFuncOneOne(fnType, elemType, nil)
}

func checkFuncOneOne(fnType, elemType, resType reflect.Type) error {
	if fnType.Kind() != reflect.Func {
		return errors.New("fn is not func")
	}

	if fnType.NumIn() != 1 {
		return errors.New("number of args is not equal 1")
	}

	if fnType.NumOut() != 1 {
		return errors.New("number of outputs is not equal 1")
	}

	if fnType.In(0) != elemType {
		return fmt.Errorf(fmt.Sprintf("arg has wrong type, actual %s, need %s", fnType.In(0), elemType))
	}

	if resType != nil {
		if fnType.Out(0) != resType {
			return fmt.Errorf("output has wrong type, actual %s, need %s", fnType.Out(0), resType)
		}
	}

	return nil
}
