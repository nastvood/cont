package fslice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFold(t *testing.T) {
	type args struct {
		s    interface{}
		zero interface{}
		fn   interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "nil",
			args: args{
				s:    nil,
				zero: 0,
				fn: func(v, acc int) int {
					return v + acc
				},
			},
			want: 0,
		},
		{
			name: "not slice",
			args: args{
				s:    1,
				zero: 0,
				fn: func(v, acc int) int {
					return v + acc
				},
			},
			want: 0,
		},
		{
			name: "check err",
			args: args{
				s:    []int64{1, 3},
				zero: 0,
				fn: func(v bool, acc int) int {
					return 1 + acc
				},
			},
			want: 0,
		},
		{
			name: "empty",
			args: args{
				s:    []int{},
				zero: 0,
				fn: func(v, acc int) int {
					return v + acc
				},
			},
			want: 0,
		},
		{
			name: "sum",
			args: args{
				s:    []int{1, 2, 3, 4},
				zero: 0,
				fn: func(v, acc int) int {
					return v + acc
				},
			},
			want: 10,
		},
		{
			name: "concat",
			args: args{
				s:    []string{"1", "2", "3", "4", "xyz"},
				zero: "",
				fn: func(s, acc string) string {
					return acc + s
				},
			},
			want: "1234xyz",
		},
		{
			name: "int64 to user",
			args: args{
				s:    []int64{1, 2, 3},
				zero: make(map[int64]string),
				fn: func(i int64, m map[int64]string) map[int64]string {
					m[i] = "u" + strconv.FormatInt(i, 10)
					return m
				},
			},
			want: map[int64]string{
				1: "u1",
				2: "u2",
				3: "u3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fold(tt.args.s, tt.args.zero, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}
