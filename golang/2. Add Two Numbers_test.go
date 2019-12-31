package golang

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				l1: NewList(2, 4, 3),
				l2: NewList(5, 6, 4),
			},
			want: NewList(7, 0, 8),
		},
		{
			name: "empty l1",
			args: args{
				l1: nil,
				l2: NewList(8),
			},
			want: NewList(8),
		},
		{
			name: "1 node of l1",
			args: args{
				l1: NewList(8),
				l2: NewList(2, 4, 3),
			},
			want: NewList(0, 5, 3),
		},
		{
			name: "all 9 of l1",
			args: args{
				l1: NewList(9, 9, 9),
				l2: NewList(1),
			},
			want: NewList(0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbersFromDiscuss(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				l1: NewList(2, 4, 3),
				l2: NewList(5, 6, 4),
			},
			want: NewList(7, 0, 8),
		},
		{
			name: "empty l1",
			args: args{
				l1: nil,
				l2: NewList(8),
			},
			want: NewList(8),
		},
		{
			name: "1 node of l1",
			args: args{
				l1: NewList(8),
				l2: NewList(2, 4, 3),
			},
			want: NewList(0, 5, 3),
		},
		{
			name: "all 9 of l1",
			args: args{
				l1: NewList(9, 9, 9),
				l2: NewList(1),
			},
			want: NewList(0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersFromDiscuss(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersFromDiscuss() = %v, want %v", got, tt.want)
			}
		})
	}
}
