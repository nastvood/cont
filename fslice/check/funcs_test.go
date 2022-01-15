package check

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	ferrors "github.com/nastvood/cont/errors"
	"github.com/nastvood/cont/internal/pkg/util"
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
			wantErr: ferrors.NewNotFuncError(),
		},
		{
			name: "num in 2",
			args: args{
				fnType: reflect.TypeOf(func(i int) int {
					return 0
				}),
			},
			wantErr: ferrors.NewNumInError(2),
		},
		{
			name: "num out 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, j int) {
				}),
			},
			wantErr: ferrors.NewNumOutError(1),
		},
		{
			name: "in type 0",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) bool {
					return false
				}),
				elemType: util.BoolType,
			},
			wantErr: ferrors.NewInTypeError(0, util.BoolType, util.IntType),
		},
		{
			name: "in type 1",
			args: args{
				fnType: reflect.TypeOf(func(i int, b bool) int {
					return i
				}),
				elemType: util.IntType,
			},
			wantErr: ferrors.NewInTypeError(1, util.BoolType, util.IntType),
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
			if err != nil && tt.wantErr != nil {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
