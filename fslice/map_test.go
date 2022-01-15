package slice

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
			name: "empty",
			args: args{
				s:  []int{},
				fn: func() {},
			},
			want: []int{},
		},
		{
			name: "int to int",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) int {
					return i * i
				},
			},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "string to int64",
			args: args{
				s: []string{"1", "2", "3", "4", "xyz"},
				fn: func(s string) int64 {
					i, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return -1
					}

					return i * 2
				},
			},
			want: []int64{2, 4, 6, 8, -1},
		},
		{
			name: "int64 to user",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int64) user {
					return user{
						id:   i,
						name: "user" + strconv.FormatInt(i, 10),
					}
				},
			},
			want: []user{{1, "user1"}, {2, "user2"}, {3, "user3"}},
		},
		{
			name: "int64 to user pointer",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int64) *user {
					if i == 3 {
						return nil
					}

					return &user{
						id:   i,
						name: "user" + strconv.FormatInt(i, 10),
					}
				},
			},
			want: []*user{{1, "user1"}, {2, "user2"}, nil},
		},
		{
			name: "int64 to int slice",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int64) []int {
					return []int{int(i)}
				},
			},
			want: [][]int{{1}, {2}, {3}},
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

func TestMapi(t *testing.T) {
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
			name: "empty",
			args: args{
				s:  []int{},
				fn: func() {},
			},
			want: []int{},
		},
		{
			name: "int to int",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i, v int) int {
					if i%2 == 0 {
						return i
					}

					return v * 2
				},
			},
			want: []int{0, 4, 2, 8},
		},
		{
			name: "string to int64",
			args: args{
				s: []string{"1", "2", "3", "4", "xyz"},
				fn: func(i int, s string) int64 {
					v, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						return -1
					}

					if i%2 == 0 {
						return v
					}

					return v * 2
				},
			},
			want: []int64{1, 4, 3, 8, -1},
		},
		{
			name: "int64 to user",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int, v int64) user {
					return user{
						id:   int64(i),
						name: "user" + strconv.FormatInt(v, 10),
					}
				},
			},
			want: []user{{0, "user1"}, {1, "user2"}, {2, "user3"}},
		},
		{
			name: "int64 to user pointer",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int, v int64) *user {
					if i%2 == 0 {
						return nil
					}

					return &user{
						id:   int64(i),
						name: "user" + strconv.FormatInt(v, 10),
					}
				},
			},
			want: []*user{nil, {1, "user2"}, nil},
		},
		{
			name: "int64 to int slice",
			args: args{
				s: []int64{1, 2, 3},
				fn: func(i int, v int64) []int {
					return []int{i, int(v)}
				},
			},
			want: [][]int{{0, 1}, {1, 2}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mapi(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mapi() = %v, want %v", got, tt.want)
			}
		})
	}
}
