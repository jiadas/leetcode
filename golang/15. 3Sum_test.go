package golang

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var sort2DIntSlice = cmp.Transformer("Sort2DIntSlice", func(in [][]int) [][]int {
	var out [][]int
	for i := range in {
		c := append([]int(nil), in[i]...)
		sort.Ints(c)
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool {
		for x := range out[i] {
			if out[i][x] == out[j][x] {
				continue
			}
			return out[i][x] < out[j][x]
		}
		return false
	})
	return out
})

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "e1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, 1, 0}, {-1, 2, -1}},
		},
		{
			name: "3",
			args: args{
				nums: []int{-1, 0, 1},
			},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "all zero",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "dup",
			args: args{
				nums: []int{1, 2, -2, -1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !cmp.Equal(got, tt.want, sort2DIntSlice) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
