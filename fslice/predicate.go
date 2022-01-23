package fslice

import (
	"reflect"

	"github.com/nastvood/cont/fslice/check"
)

// ForAll - func({s0, s1, ..., sn}, fn) checks if all elements of the slice satisfy the predicate fn.
// That is, it returns fn(s0) && fn(s1) && ... && fn(sn).
func ForAll(s, fn interface{}) bool {
	if s == nil {
		return true
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return false
	}

	elType := reflect.TypeOf(s).Elem()
	err := check.FuncUnaryPredicate(reflect.TypeOf(fn), elType)
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

// Exists - func({s0, s1, ..., sn}, fn) checks if at least one element of the slice satisfies the predicate fn.
// That is, it returns fn(s0) || fn(s1) || ... || fn(sn).
func Exists(s, fn interface{}) bool {
	if s == nil {
		return false
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return false
	}

	elType := reflect.TypeOf(s).Elem()
	err := check.FuncUnaryPredicate(reflect.TypeOf(fn), elType)
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
