package fslice

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		s    interface{}
		fn   interface{}
		opts []findOption
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil",
			args: args{
				s:  nil,
				fn: func() {},
			},
			want: -1,
		},
		{
			name: "not slice",
			args: args{
				s:  1,
				fn: func() {},
			},
			want: -1,
		},
		{
			name: "wron func",
			args: args{
				s:  []int{1, 2, 3},
				fn: func() {},
			},
			want: -1,
		},
		{
			name: "empty",
			args: args{
				s: []int{},
				fn: func(i int) bool {
					return false
				},
			},
			want: -1,
		},
		{
			name: "find",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i > 2
				},
			},
			want: 2,
		},
		{
			name: "find from right",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i > 2
				},
				opts: []findOption{WithFindFromRight()},
			},
			want: 3,
		},
		{
			name: "not found",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i > 4
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.s, tt.args.fn, tt.args.opts...); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
