package golang

import (
	"reflect"
	"testing"
)

func Test_doOddEvenList(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			name: "wa1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doOddEvenList(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doOddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
