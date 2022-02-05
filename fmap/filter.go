package fmap

import (
	"reflect"

	"github.com/barkimedes/go-deepcopy"

	"github.com/nastvood/cont/fmap/check"
)

type filterOption func(c *filterConfig)

type filterConfig struct {
	deepCopy bool
}

func WithFilterDeepCopy() filterOption {
	return func(c *filterConfig) {
		c.deepCopy = true
	}
}

// Filter - func(m, fn) returns all the elements of the map m that
// satisfy the predicate fn.
// fn(key, value) bool
func Filter(m, fn interface{}, opts ...filterOption) interface{} {
	if m == nil {
		return m
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return m
	}

	fnType := reflect.TypeOf(fn)
	mType := reflect.TypeOf(m)
	if err := check.FuncUnaryPredicate(fnType, mType.Key(), mType.Elem()); err != nil {
		return m
	}

	c := &filterConfig{}
	for _, opt := range opts {
		opt(c)
	}

	newMap := reflect.MakeMap(reflect.TypeOf(m))

	iter := val.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		ok := reflect.ValueOf(fn).Call([]reflect.Value{key, value})[0]
		if ok.Bool() {
			if c.deepCopy {
				newValue := deepcopy.MustAnything(value.Interface())
				newMap.SetMapIndex(key, reflect.ValueOf(newValue))

				continue
			}

			newMap.SetMapIndex(key, value)
		}

	}

	return newMap.Interface()
}
