package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate90(t *testing.T) {
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
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate90(tt.args.input)
			assert.Equal(t, tt.args.expect, tt.args.input)
		})
	}
}
