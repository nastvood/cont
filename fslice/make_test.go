package fslice

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/nastvood/cont/internal/pkg/util"
)

func TestMake(t *testing.T) {
	type (
		user struct {
			Id   int64
			Name string
		}

		args struct {
			n  int
			el interface{}
		}
	)
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "negative",
			args: args{
				n:  -1,
				el: nil,
			},
			wantErr: true,
		},
		{
			name: "nil",
			args: args{
				n:  5,
				el: nil,
			},
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				n: 0,
				el: func() *user {
					return nil
				}(),
			},
			want: []*user{},
		},
		{
			name: "nil user",
			args: args{
				n: 2,
				el: func() *user {
					return nil
				}(),
			},
			want: []*user{nil, nil},
		},
		{
			name: "pointer user",
			args: args{
				n: 2,
				el: func() *user {
					return &user{
						Id:   1,
						Name: "user1",
					}
				}(),
			},
			want: []*user{{1, "user1"}, {1, "user1"}},
		},
		{
			name: "user",
			args: args{
				n: 1,
				el: func() user {
					return user{
						Id:   1,
						Name: "user1",
					}
				}(),
			},
			want: []user{{1, "user1"}},
		},
		{
			name: "nil int",
			args: args{
				n: 3,
				el: func() *int {
					return nil
				}(),
			},
			want: []*int{nil, nil, nil},
		},
		{
			name: "pointer int",
			args: args{
				n: 3,
				el: func() *int {
					i := 1
					return &i
				}(),
			},
			want: []*int{util.NewInt(1), util.NewInt(1), util.NewInt(1)},
		},
		{
			name: "int",
			args: args{
				n: 2,
				el: func() int {
					return 1
				}(),
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Make(tt.args.n, tt.args.el)
			if (err != nil) != tt.wantErr {
				t.Errorf("Make() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	type (
		user struct {
			Id   int64
			Name string
		}

		args struct {
			n  int
			fn interface{}
		}
	)
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "negative",
			args: args{
				n:  -1,
				fn: func() {},
			},
			wantErr: true,
		},
		{
			name: "wrong func",
			args: args{
				n:  1,
				fn: func() {},
			},
			wantErr: true,
		},
		{
			name: "empty",
			args: args{
				n: 0,
				fn: func(i int) *user {
					return nil
				},
			},
			want: []*user{},
		},
		{
			name: "nil user",
			args: args{
				n: 2,
				fn: func(i int) *user {
					return nil
				},
			},
			want: []*user{nil, nil},
		},
		{
			name: "pointer user",
			args: args{
				n: 2,
				fn: func(i int) *user {
					return &user{
						Id:   int64(i + 1),
						Name: "user" + strconv.FormatInt(int64(i+1), 10),
					}
				},
			},
			want: []*user{{1, "user1"}, {2, "user2"}},
		},
		{
			name: "user",
			args: args{
				n: 1,
				fn: func(i int) user {
					return user{
						Id:   int64(i + 1),
						Name: "user" + strconv.FormatInt(int64(i+1), 10),
					}
				},
			},
			want: []user{{1, "user1"}},
		},
		{
			name: "nil int",
			args: args{
				n: 3,
				fn: func(i int) *int {
					return nil
				},
			},
			want: []*int{nil, nil, nil},
		},
		{
			name: "pointer int",
			args: args{
				n: 3,
				fn: func(i int) *int {
					v := i + 1
					return &v
				},
			},
			want: []*int{util.NewInt(1), util.NewInt(2), util.NewInt(3)},
		},
		{
			name: "int",
			args: args{
				n: 2,
				fn: func(i int) int {
					return i
				},
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.n, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
