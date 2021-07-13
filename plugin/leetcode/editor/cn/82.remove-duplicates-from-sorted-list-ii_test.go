package golang

import (
	"reflect"
	"testing"
)

func Test_doDeleteDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "wa1",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 4, 5},
			},
			want: []int{1, 2, 5},
		},
		{
			name: "wa2",
			args: args{
				nums: []int{1, 2, 2},
			},
			want: []int{1},
		},
		{
			name: "wa3",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doDeleteDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doDeleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
