package fmap

import (
	"reflect"

	"github.com/barkimedes/go-deepcopy"
)

type valuesOption func(*valuesConfig)

type valuesConfig struct {
	deepCopy bool
}

func WithValuesDeepCopy() valuesOption {
	return func(config *valuesConfig) {
		config.deepCopy = true
	}
}

func Values(m interface{}, opts ...valuesOption) interface{} {
	if m == nil {
		return m
	}

	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return m
	}

	values := reflect.MakeSlice(reflect.SliceOf(val.Type().Elem()), 0, val.Len())
	if val.Len() == 0 {
		return values.Interface()
	}

	config := &valuesConfig{}
	for _, opt := range opts {
		opt(config)
	}

	iter := val.MapRange()
	for iter.Next() {
		if config.deepCopy {
			newEl := deepcopy.MustAnything(iter.Key().Interface())
			values = reflect.Append(values, reflect.ValueOf(newEl))
			continue
		}

		values = reflect.Append(values, iter.Value())
	}

	return values.Interface()
}
