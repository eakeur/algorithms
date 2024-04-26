package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteIndex(t *testing.T) {
	type args struct {
		a []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				a: []int{1, 2, 3},
				i: 1,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				i: 3,
			},
			want: []int{1, 2, 3, 5, 6, 7},
		},
		{
			args: args{
				a: []int{1},
				i: 1,
			},
			want: []int{1},
		},
		{
			args: args{
				a: []int{1},
				i: 0,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.want, deleteIndex(tt.args.a, tt.args.i))
		})
	}
}

func Test_deleteIndexInPlace(t *testing.T) {
	type args struct {
		a []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				a: []int{1, 2, 3},
				i: 1,
			},
			want: []int{1, 3, 0},
		},
		{
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 7},
				i: 3,
			},
			want: []int{1, 2, 3, 5, 6, 7, 0},
		},
		{
			args: args{
				a: []int{1},
				i: 1,
			},
			want: []int{1},
		},
		{
			args: args{
				a: []int{1},
				i: 0,
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteIndexInPlace(tt.args.a, tt.args.i)
			assert.Equal(t, tt.want, tt.args.a)
		})
	}
}
