package array

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{4, 2, 1, 3},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			args: args{
				arr: []int{10, 32, 0, -55, 44, 22, 0, 2, 1, 2, 5, 8, 9},
			},
			want: []int{-55, 0, 0, 1, 2, 2, 5, 8, 9, 10, 22, 32, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
