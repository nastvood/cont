package check

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	cerrors "github.com/nastvood/cont/errors"
	"github.com/nastvood/cont/internal/util"
)

func TestFuncMap(t *testing.T) {
	type args struct {
		fnType    reflect.Type
		keyType   reflect.Type
		valueType reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "not func",
			args: args{
				fnType: util.IntType,
			},
			wantErr: cerrors.NewNotFuncError(),
		},
		{
			name: "num in 2",
			args: args{
				fnType: reflect.TypeOf(func(i int) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(2),
		},
		{
			name: "num out 2",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumOutError(2),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) (int, bool) {
					return 0, true
				}),
				keyType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) (int, bool) {
					return k, v
				}),
				keyType:   util.IntType,
				valueType: util.IntType,
			},
			wantErr: cerrors.NewInTypeError(1, util.IntType, util.BoolType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) (int, bool) {
					return k, v
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncMap(tt.args.fnType, tt.args.keyType, tt.args.valueType)
			if err != nil && tt.wantErr == nil ||
				err == nil && tt.wantErr != nil {
				require.Failf(t, "err", "err is nil (%t), wantErr is nil (%t)", err == nil, tt.wantErr == nil)
			}

			if err != nil || tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestFuncFold(t *testing.T) {
	type args struct {
		fnType    reflect.Type
		keyType   reflect.Type
		valueType reflect.Type
		zeroType  reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "not func",
			args: args{
				fnType: util.IntType,
			},
			wantErr: cerrors.NewNotFuncError(),
		},
		{
			name: "num in 3",
			args: args{
				fnType: reflect.TypeOf(func(i int) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(3),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool, acc int) (int, error) {
					return 0, nil
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool, acc float64) float64 {
					return 0.0
				}),
				keyType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool, acc float64) float64 {
					return 0.0
				}),
				keyType:   util.IntType,
				valueType: util.IntType,
			},
			wantErr: cerrors.NewInTypeError(1, util.IntType, util.BoolType),
		},
		{
			name: "out type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool, acc float64) string {
					return ""
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
				zeroType:  util.Float64Type,
			},
			wantErr: cerrors.NewOutTypeError(0, util.Float64Type, util.StringType),
		},
		{
			name: "out type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool, acc float64) float64 {
					return 0.0
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
				zeroType:  util.StringType,
			},
			wantErr: cerrors.NewOutTypeError(0, util.StringType, util.Float64Type),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool, acc float64) float64 {
					return 0.0
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
				zeroType:  util.Float64Type,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncFold(tt.args.fnType, tt.args.keyType, tt.args.valueType, tt.args.zeroType)
			if err != nil && tt.wantErr == nil ||
				err == nil && tt.wantErr != nil {
				require.Failf(t, "err", "err is nil (%t), wantErr is nil (%t)", err == nil, tt.wantErr == nil)
			}

			if err != nil || tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestFuncUnaryPredicate(t *testing.T) {
	type args struct {
		fnType    reflect.Type
		keyType   reflect.Type
		valueType reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "not func",
			args: args{
				fnType: util.IntType,
			},
			wantErr: cerrors.NewNotFuncError(),
		},
		{
			name: "num in 2",
			args: args{
				fnType: reflect.TypeOf(func(i int) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(2),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) (int, error) {
					return 0, nil
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) bool {
					return true
				}),
				keyType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) bool {
					return true
				}),
				keyType:   util.IntType,
				valueType: util.IntType,
			},
			wantErr: cerrors.NewInTypeError(1, util.IntType, util.BoolType),
		},
		{
			name: "out type 0",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) string {
					return ""
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
			},
			wantErr: cerrors.NewOutTypeError(0, util.BoolType, util.StringType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(k int, v bool) bool {
					return true
				}),
				keyType:   util.IntType,
				valueType: util.BoolType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncUnaryPredicate(tt.args.fnType, tt.args.keyType, tt.args.valueType)
			if err != nil && tt.wantErr == nil ||
				err == nil && tt.wantErr != nil {
				require.Failf(t, "err", "err is nil (%t), wantErr is nil (%t)", err == nil, tt.wantErr == nil)
			}

			if err != nil || tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
