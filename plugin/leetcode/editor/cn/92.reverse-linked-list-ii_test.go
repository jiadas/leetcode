package golang

import (
	"reflect"
	"testing"
)

func Test_doReverseBetween(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "e1",
			args: args{
				nums:  []int{1, 2, 3, 4, 5},
				left:  2,
				right: 4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "e2",
			args: args{
				nums:  []int{5},
				left:  1,
				right: 1,
			},
			want: []int{5},
		},
		{
			name: "e3",
			args: args{
				nums:  []int{1, 2, 3, 4, 5, 6},
				left:  1,
				right: 5,
			},
			want: []int{5, 4, 3, 2, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doReverseBetween(tt.args.nums, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doReverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
