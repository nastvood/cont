package errors

import (
	"fmt"
	"reflect"
	"strconv"
)

type NotFuncError struct{}

func (e *NotFuncError) Error() string {
	return "not func"
}

func NewNotFuncError() *NotFuncError {
	return &NotFuncError{}
}

type NumInError struct {
	Num int
}

func (e *NumInError) Error() string {
	return "number of args is not equal " + strconv.FormatInt(int64(e.Num), 10)
}

func NewNumInError(num int) *NumInError {
	return &NumInError{
		Num: num,
	}
}

type NumOutError struct {
	Num int
}

func (e *NumOutError) Error() string {
	return "number of outputs is not equal " + strconv.FormatInt(int64(e.Num), 10)
}

func NewNumOutError(num int) *NumOutError {
	return &NumOutError{
		Num: num,
	}
}

type InTypeError struct {
	Index        int
	ExpectedType reflect.Type
	ActualType   reflect.Type
}

func (e *InTypeError) Error() string {
	return fmt.Sprintf("arg %d has wrong type, actual %s, expected %s", e.Index, e.ActualType, e.ExpectedType)
}

func NewInTypeError(index int, expectedType reflect.Type, actualType reflect.Type) *InTypeError {
	return &InTypeError{
		Index:        index,
		ExpectedType: expectedType,
		ActualType:   actualType,
	}
}

type OutTypeError struct {
	Index        int
	ExpectedType reflect.Type
	ActualType   reflect.Type
}

func (e *OutTypeError) Error() string {
	return fmt.Sprintf("output %d has wrong type, actual %s, expected %s", e.Index, e.ActualType, e.ExpectedType)
}

func NewOutTypeError(index int, expectedType reflect.Type, actualType reflect.Type) *OutTypeError {
	return &OutTypeError{
		Index:        index,
		ExpectedType: expectedType,
		ActualType:   actualType,
	}
}
