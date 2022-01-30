package fmap

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValues(t *testing.T) {
	type (
		user = struct {
			ID     int
			Emails []string
		}

		args struct {
			m    interface{}
			opts []valuesOption
		}
	)
	tests := []struct {
		name     string
		args     args
		lessFunc func(s interface{}) func(i, j int) bool
		want     interface{}
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
			want: []string{},
		},
		{
			name: "succ",
			args: args{
				m: map[int]string{
					1: "1",
					2: "2",
				},
			},
			want: []string{"1", "2"},
			lessFunc: func(got interface{}) func(i, j int) bool {
				return func(i, j int) bool {
					s, ok := got.([]string)
					require.True(t, ok)
					return s[i] < s[j]
				}
			},
		},
		{
			name: "succ with option",
			args: args{
				m: map[int]user{
					1: {
						ID:     1,
						Emails: []string{"e1_1@example.com", "e1_2@example.com"},
					},
					2: {
						ID:     2,
						Emails: []string{"e2_1@example.com", "e2_2@example.com"},
					},
				},
				opts: []valuesOption{WithValuesDeepCopy()},
			},
			want: []user{{
				ID:     1,
				Emails: []string{"e1_1@example.com", "e1_2@example.com"},
			}, {
				ID:     2,
				Emails: []string{"e2_1@example.com", "e2_2@example.com"},
			}},
			lessFunc: func(got interface{}) func(i, j int) bool {
				return func(i, j int) bool {
					s, ok := got.([]user)
					require.True(t, ok)
					return s[i].ID < s[j].ID
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Values(tt.args.m, tt.args.opts...)
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
