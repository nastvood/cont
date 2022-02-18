package check

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	cerrors "github.com/nastvood/cont/errors"
	"github.com/nastvood/cont/internal/util"
)

func TestFuncFold(t *testing.T) {
	type args struct {
		fnType   reflect.Type
		elemType reflect.Type
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
				fnType: reflect.TypeOf(func(i int, j int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return i
				}),
				elemType: util.IntType,
			},
			wantErr: cerrors.NewInTypeError(1, util.BoolType, util.IntType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) bool {
					return false
				}),
				elemType: util.IntType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncFold(tt.args.fnType, tt.args.elemType)
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

func TestFuncMapi(t *testing.T) {
	type args struct {
		fnType   reflect.Type
		elemType reflect.Type
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
				fnType: reflect.TypeOf(func(i int, j int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(el bool, b bool) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.IntType, util.BoolType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return i
				}),
				elemType: util.IntType,
			},
			wantErr: cerrors.NewInTypeError(1, util.IntType, util.BoolType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncMapi(tt.args.fnType, tt.args.elemType)
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

func TestFuncFilterMap(t *testing.T) {
	type args struct {
		fnType   reflect.Type
		elemType reflect.Type
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
			name: "num in 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(1),
		},
		{
			name: "num out 2",
			args: args{
				fnType: reflect.TypeOf(func(i int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(2),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(el int) (int, bool) {
					return 0, false
				}),
				elemType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "out type 1",
			args: args{
				fnType: reflect.TypeOf(func(i int) (int, int) {
					return i, 0
				}),
				elemType: util.IntType,
			},
			wantErr: cerrors.NewOutTypeError(1, util.BoolType, util.IntType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int) (float64, bool) {
					return .0, false
				}),
				elemType: util.IntType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncFilterMap(tt.args.fnType, tt.args.elemType)
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

func TestFuncInit(t *testing.T) {
	type args struct {
		fnType reflect.Type
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
			name: "num in 1",
			args: args{
				fnType: reflect.TypeOf(func() int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(1),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(s string) bool {
					return false
				}),
			},
			wantErr: cerrors.NewInTypeError(0, util.IntType, util.StringType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int) bool {
					return i > 0
				}),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncInit(tt.args.fnType)
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
		fnType   reflect.Type
		elemType reflect.Type
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
			name: "num in 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(1),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(el int) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "out type 0",
			args: args{
				fnType: reflect.TypeOf(func(i int) int {
					return i
				}),
				elemType: util.IntType,
			},
			wantErr: cerrors.NewOutTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int) bool {
					return false
				}),
				elemType: util.IntType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncUnaryPredicate(tt.args.fnType, tt.args.elemType)
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

func TestFuncMap(t *testing.T) {
	type args struct {
		fnType   reflect.Type
		elemType reflect.Type
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
			name: "num in 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return 0
				}),
			},
			wantErr: cerrors.NewNumInError(1),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int) {
				}),
			},
			wantErr: cerrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(el int) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: cerrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "succ",
			args: args{
				fnType: reflect.TypeOf(func(i int) bool {
					return false
				}),
				elemType: util.IntType,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FuncMap(tt.args.fnType, tt.args.elemType)
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
