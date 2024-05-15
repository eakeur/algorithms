package linkedlist

import (
	"reflect"
	"testing"
)

func Test_node(t *testing.T) {

	type args struct {
		val int
		del int
	}
	tests := []struct {
		name    string
		args    args
		want    *node
		wantDel *node
	}{
		{
			args: args{
				val: 3,
				del: 1,
			},
			want: &node{
				value: 0,
				next: &node{
					value: 1,
					next: &node{
						value: 2,
					},
				},
			},
			wantDel: &node{
				value: 0,
				next: &node{
					value: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n *node

			for i := 0; i < tt.args.val; i++ {
				n = n.add(i)
			}

			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("node.add() = %v, want %v", n, tt.want)
			}

			n = n.delete(tt.args.del)

			if !reflect.DeepEqual(n, tt.wantDel) {
				t.Errorf("node.add() = %v, want %v", n, tt.wantDel)
			}
		})
	}
}
