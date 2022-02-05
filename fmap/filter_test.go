package fmap

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type (
		args struct {
			m    interface{}
			fn   interface{}
			opts []filterOption
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
				m: nil,
				fn: func() int {
					return 0
				},
			},
			want: nil,
		},
		{
			name: "not map",
			args: args{
				m: 1,
				fn: func() int {
					return 0
				},
			},
			want: 1,
		},
		{
			name: "check err",
			args: args{
				m: map[string]int64{"1": 1},
				fn: func(v bool, acc int) int {
					return 1 + acc
				},
			},
			want: map[string]int64{"1": 1},
		},
		{
			name: "gr",
			args: args{
				m: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4},
				fn: func(k string, v int) bool {
					return v > 3
				},
			},
			want: map[string]int{"4": 4},
		},
		{
			name: "deep copy",
			args: args{
				m: map[string][]int{
					"1": {11, 12, 13},
					"2": {21, 22, 23},
					"3": {31, 32, 33},
				},
				fn: func(k string, v []int) bool {
					return k == "2" || k == "3"
				},
				opts: []filterOption{
					WithFilterDeepCopy(),
				},
			},
			want: map[string][]int{
				"2": {21, 22, 23},
				"3": {31, 32, 33},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.m, tt.args.fn, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
