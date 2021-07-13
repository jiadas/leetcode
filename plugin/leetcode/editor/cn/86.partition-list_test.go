package golang

import (
	"reflect"
	"testing"
)

func Test_doPartition(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{1, 4, 3, 2, 5, 2},
				x:    3,
			},
			want: []int{1, 2, 2, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPartition(tt.args.nums, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
