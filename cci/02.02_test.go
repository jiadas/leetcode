package cci

import "testing"

func Test_kthToLast(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				k: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.head, tt.args.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
