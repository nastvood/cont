package fmap

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type (
		user struct {
			id   int64
			name string
		}

		args struct {
			s  interface{}
			fn interface{}
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
				s:  nil,
				fn: func() {},
			},
			want: nil,
		},
		{
			name: "not map",
			args: args{
				s:  1,
				fn: func() {},
			},
			want: 1,
		},
		{
			name: "check func err",
			args: args{
				s:  map[int]int{},
				fn: func() {},
			},
			want: map[int]int{},
		},
		{
			name: "int to int",
			args: args{
				s: map[int]int{
					1: 10,
					2: 20,
					3: 30,
				},
				fn: func(k int, v int) (int, int) {
					return k * v, v * v
				},
			},
			want: map[int]int{10: 100, 40: 400, 90: 900},
		},
		{
			name: "string to int64",
			args: args{
				s: map[string]int64{"1": 1, "2": 2, "xyz": 0},
				fn: func(k string, v int64) (int64, string) {
					i, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						return -1, k
					}

					return i * v * 2, k
				},
			},
			want: map[int64]string{-1: "xyz", 2: "1", 8: "2"},
		},
		{
			name: "int64 to user",
			args: args{
				s: map[int64]struct{}{
					0: {},
					1: {},
				},
				fn: func(k int64, v struct{}) (int64, user) {
					return k + 1, user{
						id:   k,
						name: "user" + strconv.FormatInt(k, 10),
					}
				},
			},
			want: map[int64]user{1: {0, "user0"}, 2: {1, "user1"}},
		},
		{
			name: "int64 to user pointer",
			args: args{
				s: map[int64]struct{}{
					0: {},
					1: {},
					2: {},
				},
				fn: func(k int64, v struct{}) (int64, *user) {
					if k == 2 {
						return 2, nil
					}

					return k, &user{
						id:   k,
						name: "user" + strconv.FormatInt(k, 10),
					}
				},
			},
			want: map[int64]*user{0: {0, "user0"}, 1: {1, "user1"}, 2: nil},
		},
		{
			name: "int to int slice",
			args: args{
				s: map[int]struct{}{
					0:  {},
					1:  {},
					2:  {},
					-2: {},
				},
				fn: func(k int, v struct{}) (int, []int) {
					l := k
					if k < 0 {
						l = 0
					}

					s := make([]int, l)
					for i := 0; i < l; i++ {
						s[i] = i + k
					}

					return k, s
				},
			},
			want: map[int][]int{-2: {}, 0: {}, 1: {1}, 2: {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
