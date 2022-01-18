package fmap

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFold(t *testing.T) {
	type (
		args struct {
			m    interface{}
			zero interface{}
			fn   interface{}
		}

		user struct {
			id   int64
			name string
		}
	)
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "nil",
			args: args{
				m:    nil,
				zero: 0,
				fn: func(v, acc int) int {
					return v + acc
				},
			},
			want: 0,
		},
		{
			name: "not map",
			args: args{
				m:    1,
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
				m:    map[string]int64{"1": 1, "2": 3},
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
				m:    map[string]int64{},
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
				m:    map[string]int{"1": 1, "2": 2, "3": 3, "4": 4},
				zero: 0,
				fn: func(k string, v int, acc int) int {
					return v + acc
				},
			},
			want: 10,
		},
		{
			name: "concat",
			args: args{
				m:    map[string]struct{}{"1": {}, "2": {}, "3": {}, "4": {}, "xyz": {}},
				zero: int64(5),
				fn: func(k string, v struct{}, acc int64) int64 {
					i, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						return acc
					}
					return acc + i
				},
			},
			want: int64(15),
		},
		{
			name: "int64 to user",
			args: args{
				m:    map[int64]string{1: "name1", 2: "name2", 3: "name3"},
				zero: make(map[int64]user),
				fn: func(k int64, v string, m map[int64]user) map[int64]user {
					m[k] = user{
						id:   k,
						name: v,
					}

					return m
				},
			},
			want: map[int64]user{
				1: {1, "name1"},
				2: {2, "name2"},
				3: {3, "name3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fold(tt.args.m, tt.args.zero, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}
