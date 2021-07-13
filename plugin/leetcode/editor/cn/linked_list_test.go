package golang

import (
	"reflect"
	"testing"
)

func Test_roundConvert(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "intSlice<->listNode",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundConvertForTest(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("roundConvertForTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
