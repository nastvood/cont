package fslice

import (
	"reflect"

	"github.com/barkimedes/go-deepcopy"

	"github.com/nastvood/cont/fslice/check"
)

type filterOption func(*filterConfig)

type filterConfig struct {
	deepCopy bool
}

func WithFilterDeepCopy() filterOption {
	return func(c *filterConfig) {
		c.deepCopy = true
	}
}

// Filter - func({s0, s1, ..., sn}, fn) returns all the elements of the slice s that
// satisfy the predicate fn.
func Filter(s, fn interface{}, opts ...filterOption) interface{} {
	if s == nil {
		return s
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return s
	}

	elType := reflect.TypeOf(s).Elem()
	err := check.FuncUnaryPredicate(reflect.TypeOf(fn), elType)
	if err != nil {
		return s
	}

	if val.Len() == 0 {
		return s
	}

	config := &filterConfig{}
	for _, opt := range opts {
		opt(config)
	}

	newS := reflect.MakeSlice(reflect.SliceOf(elType), 0, 0)
	for i := 0; i < val.Len(); i++ {
		v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
		if v.Bool() {
			if config.deepCopy {
				newEl := deepcopy.MustAnything(val.Index(i).Interface())
				newS = reflect.Append(newS, reflect.ValueOf(newEl))
				continue
			}

			newS = reflect.Append(newS, val.Index(i))
		}
	}

	return newS.Interface()
}

// FilterMap - func({s0, s1, ..., sn}, fn) calls fn(s0), fn(s1), ..., fn(sn). It returns the slice of elements bi
// such as b1, ok := fn(si) (when ok is false, the corresponding element of s is discarded).
func FilterMap(s, fn interface{}) interface{} {
	if s == nil {
		return s
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return s
	}

	fnType := reflect.TypeOf(fn)
	err := check.FuncFilterMap(fnType, reflect.TypeOf(s).Elem())
	if err != nil {
		return s
	}

	newS := reflect.MakeSlice(reflect.SliceOf(fnType.Out(0)), 0, 0)
	for i := 0; i < val.Len(); i++ {
		res := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})
		if res[1].Bool() {
			newS = reflect.Append(newS, res[0])
		}
	}

	return newS.Interface()
}
