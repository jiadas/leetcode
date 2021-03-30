package golang

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "e1",
			args: args{
				nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 4, 2, 7, 5, 3, 8, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
