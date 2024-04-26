package array

import (
	"reflect"
	"testing"
)

func Test_getCircleSlice(t *testing.T) {
	type args struct {
		a     []int
		s     int
		start int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				a: []int{1, 2, 3, 4, 5},
				s: 2,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				a: []int{1, 2, 3, 4, 5},
				s: 10,
			},
			want: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
		{
			args: args{
				a:     []int{1, 2, 3, 4, 5},
				s:     10,
				start: 4,
			},
			want: []int{5, 1, 2, 3, 4, 5, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCycle(tt.args.a, tt.args.start, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCircleSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
