package golang

import (
	"reflect"
	"testing"
)

func Test_doReorderList(t *testing.T) {
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
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doReorderList(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doReorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
