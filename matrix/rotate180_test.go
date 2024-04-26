package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate180(t *testing.T) {
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
					{9, 8, 7},
					{6, 5, 4},
					{3, 2, 1},
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
					{16, 15, 14, 13},
					{12, 11, 10, 9},
					{8, 7, 6, 5},
					{4, 3, 2, 1},
				},
			},
		},
		{
			args: args{
				input: [][]int{
					{1, 2, 3, 4, 0},
					{5, 6, 7, 8, 77},
					{9, 10, 11, 12, 65},
					{13, 14, 15, 16, 45},
					{13, 4, 7, 16, 45},
				},
				expect: [][]int{
					{45, 16, 7, 4, 13},
					{45, 16, 15, 14, 13},
					{65, 12, 11, 10, 9},
					{77, 8, 7, 6, 5},
					{0, 4, 3, 2, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate180(tt.args.input)
			assert.Equal(t, tt.args.expect, tt.args.input)
		})
	}
}
