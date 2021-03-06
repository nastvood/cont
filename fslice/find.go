package fslice

import (
	"reflect"

	"github.com/nastvood/cont/fslice/check"
)

type findOption func(*findConfig)

type findConfig struct {
	fromRight bool
}

func WithFindFromRight() findOption {
	return func(c *findConfig) {
		c.fromRight = true
	}
}

// Find - fun({s0, s1, ..., sn}, fn) returns the index of first element of the slice s
// that satisfies the predicate fn. If element not found returns -1.
func Find(s, fn interface{}, opts ...findOption) int {
	if s == nil {
		return -1
	}

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return -1
	}

	elType := reflect.TypeOf(s).Elem()
	err := check.FuncUnaryPredicate(reflect.TypeOf(fn), elType)
	if err != nil {
		return -1
	}

	if val.Len() == 0 {
		return -1
	}

	config := &findConfig{}
	for _, opt := range opts {
		opt(config)
	}

	if config.fromRight {
		for i := val.Len() - 1; i >= 0; i-- {
			v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
			if v.Bool() {
				return i
			}
		}
	} else {
		for i := 0; i < val.Len(); i++ {
			v := reflect.ValueOf(fn).Call([]reflect.Value{val.Index(i)})[0]
			if v.Bool() {
				return i
			}
		}
	}

	return -1
}
