package fmap

import (
	"reflect"

	"github.com/barkimedes/go-deepcopy"
)

type keysOption func(*keysConfig)

type keysConfig struct {
	deepCopy bool
}

func WithKeysDeepCopy() keysOption {
	return func(config *keysConfig) {
		config.deepCopy = true
	}
}

func Keys(m interface{}, opts ...keysOption) interface{} {
	if m == nil {
		return m
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return m
	}

	keys := reflect.MakeSlice(reflect.SliceOf(val.Type().Key()), 0, val.Len())
	if val.Len() == 0 {
		return keys.Interface()
	}

	config := &keysConfig{}
	for _, opt := range opts {
		opt(config)
	}

	iter := val.MapRange()
	for iter.Next() {
		if config.deepCopy {
			newEl := deepcopy.MustAnything(iter.Key().Interface())
			keys = reflect.Append(keys, reflect.ValueOf(newEl))
			continue
		}

		keys = reflect.Append(keys, iter.Key())
	}

	return keys.Interface()
}
