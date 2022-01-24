package fslice

import (
	"testing"
)

func TestForAll(t *testing.T) {
	type args struct {
		s  interface{}
		fn interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				s:  nil,
				fn: func() {},
			},
			want: true,
		},
		{
			name: "not slice",
			args: args{
				s:  1,
				fn: func() {},
			},
			want: false,
		},
		{
			name: "wron func",
			args: args{
				s:  []int{1, 2, 3},
				fn: func() {},
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				s: []int{},
				fn: func(i int) bool {
					return true
				},
			},
			want: true,
		},
		{
			name: "int to false",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i%4 == 0
				},
			},
			want: false,
		},
		{
			name: "int to true",
			args: args{
				s: []int{2, 4, 6},
				fn: func(i int) bool {
					return i%2 == 0
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForAll(tt.args.s, tt.args.fn); got != tt.want {
				t.Errorf("ForAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		s  interface{}
		fn interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				s:  nil,
				fn: func() {},
			},
			want: false,
		},
		{
			name: "not slice",
			args: args{
				s:  1,
				fn: func() {},
			},
			want: false,
		},
		{
			name: "wron func",
			args: args{
				s:  []int{1, 2, 3},
				fn: func() {},
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				s: []int{},
				fn: func(i int) bool {
					return true
				},
			},
			want: false,
		},
		{
			name: "int to false",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i%5 == 0
				},
			},
			want: false,
		},
		{
			name: "int to true",
			args: args{
				s: []int{2, 4, 6},
				fn: func(i int) bool {
					return i%4 == 0
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.s, tt.args.fn); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
