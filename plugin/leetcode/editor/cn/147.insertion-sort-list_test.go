package golang

import (
	"reflect"
	"testing"
)

func Test_doInsertionSortList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "e1",
			args: args{
				nums: []int{-1, 5, 3, 4, 0},
			},
			want: []int{-1, 0, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doInsertionSortList(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doInsertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
