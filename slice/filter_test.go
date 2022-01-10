package slice

import (
	"reflect"
	"testing"

	"github.com/nastvood/cont/internal/pkg/util"
)

func TestFilter(t *testing.T) {
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
			name: "int",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i > 2
				},
			},
			want: []int{3, 4},
		},
		{
			name: "user",
			args: args{
				s: []user{{1, "user1"}, {2, "user2"}, {3, "user3"}},
				fn: func(u user) bool {
					return u.id > 2
				},
			},
			want: []user{{3, "user3"}},
		},
		{
			name: "user pointer",
			args: args{
				s: []*user{{1, "user1"}, nil, {3, "user3"}},
				fn: func(u *user) bool {
					if u == nil {
						return false
					}

					return u.id > 2 && len(u.name) > 0
				},
			},
			want: []*user{{3, "user3"}},
		},
		{
			name: "slice",
			args: args{
				s: [][]int{{1, 2}, nil, {3}, {3, 4}},
				fn: func(s []int) bool {
					return len(s) > 1
				},
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_DeepCopy(t *testing.T) {
	type (
		user struct {
			Id   int64
			Name string
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
			name: "int",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) bool {
					return i > 2
				},
			},
			want: []int{3, 4},
		},
		{
			name: "user",
			args: args{
				s: []user{{1, "user1"}, {2, "user2"}, {3, "user3"}},
				fn: func(u user) bool {
					return u.Id > 2
				},
			},
			want: []user{{3, "user3"}},
		},
		{
			name: "user pointer",
			args: args{
				s: []*user{{1, "user1"}, nil, {3, "user3"}},
				fn: func(u *user) bool {
					if u == nil {
						return false
					}

					return u.Id > 2 && len(u.Name) > 0
				},
			},
			want: []*user{{3, "user3"}},
		},
		{
			name: "slice",
			args: args{
				s: [][]int{{1, 2}, nil, {3}, {3, 4}},
				fn: func(s []int) bool {
					return len(s) > 1
				},
			},
			want: [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.fn, WithFilterDeepCopy()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
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
			name: "int",
			args: args{
				s: []int{1, 2, 3, 4},
				fn: func(i int) (float64, bool) {
					if i > 2 {
						return float64(i), true
					}

					return 0., false
				},
			},
			want: []float64{3, 4},
		},
		{
			name: "user",
			args: args{
				s: []user{{1, "user1"}, {2, "user2"}, {3, "user3"}},
				fn: func(u user) (int, bool) {
					if u.id > 2 {
						return int(u.id), true
					}

					return 0, false
				},
			},
			want: []int{3},
		},
		{
			name: "user pointer",
			args: args{
				s: []*user{{1, "user1"}, nil, {3, "user3"}},
				fn: func(u *user) (*string, bool) {
					if u == nil {
						return nil, false
					}

					return &u.name, true
				},
			},
			want: []*string{util.NewString("user1"), util.NewString("user3")},
		},
		{
			name: "slice",
			args: args{
				s: [][]int{{1, 2}, nil, {3}, {3, 4}},
				fn: func(s []int) ([]int, bool) {
					if len(s) == 0 {
						return nil, false
					}

					return s, true
				},
			},
			want: [][]int{{1, 2}, {3}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterMap(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
