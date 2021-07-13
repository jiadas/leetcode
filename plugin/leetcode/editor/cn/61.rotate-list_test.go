package golang

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	one := &ListNode{
		Val:  1,
		Next: nil,
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "k=1",
			args: args{
				head: one,
				k:    1,
			},
			want: one,
		},
		{
			name: "k=99",
			args: args{
				head: one,
				k:    99,
			},
			want: one,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
