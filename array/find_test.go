package array

import (
	"testing"
)

func Test_findInSortedArray(t *testing.T) {
	type args struct {
		a   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 16,
			},
			want: 15,
		},
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 22,
			},
			want: -1,
		},
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInSortedArray(tt.args.a, tt.args.num); got != tt.want {
				t.Errorf("findInSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bruteForceSearch(t *testing.T) {
	type args struct {
		a   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 16,
			},
			want: 15,
		},
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 22,
			},
			want: -1,
		},
		{
			args: args{
				a:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				num: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bruteForceSearch(tt.args.a, tt.args.num); got != tt.want {
				t.Errorf("bruteForceSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
