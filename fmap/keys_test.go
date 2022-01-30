package fmap

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	type args struct {
		m interface{}
	}

	tests := []struct {
		name     string
		args     args
		want     interface{}
		lessFunc func(s interface{}) func(i, j int) bool
	}{
		{
			name: "nil",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "not map",
			args: args{
				m: 5,
			},
			want: 5,
		},
		{
			name: "empty",
			args: args{
				m: map[int]string{},
			},
			want: []int{},
			lessFunc: func(got interface{}) func(i, j int) bool {
				return func(i, j int) bool {
					s, ok := got.([]int)
					require.True(t, ok)
					return s[i] < s[j]
				}
			},
		},
		{
			name: "succ",
			args: args{
				m: map[int]string{
					1: "1",
					2: "2",
				},
			},
			want: []int{1, 2},
			lessFunc: func(got interface{}) func(i, j int) bool {
				return func(i, j int) bool {
					s, ok := got.([]int)
					require.True(t, ok)
					return s[i] < s[j]
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys(tt.args.m)
			val := reflect.ValueOf(got)
			if val.Kind() != reflect.Slice {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Values() = %#v, want %#v", got, tt.want)
				}

				return
			}

			if tt.lessFunc != nil {
				sort.Slice(got, tt.lessFunc(got))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
