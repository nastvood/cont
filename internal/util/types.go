package util

import "reflect"

var (
	BoolType    = reflect.TypeOf(false)
	IntType     = reflect.TypeOf(int(0))
	StringType  = reflect.TypeOf("")
	Float64Type = reflect.TypeOf(0.)
)
