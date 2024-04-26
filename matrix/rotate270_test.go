package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate270(t *testing.T) {
	type args struct {
		input  [][]int
		expect [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				input: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				expect: [][]int{
					{3, 6, 9},
					{2, 5, 8},
					{1, 4, 7},
				},
			},
		},
		{
			args: args{
				input: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
				expect: [][]int{
					{4, 8, 12, 16},
					{3, 7, 11, 15},
					{2, 6, 10, 14},
					{1, 5, 9, 13},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate270(tt.args.input)
			assert.Equal(t, tt.args.expect, tt.args.input)
		})
	}
}
