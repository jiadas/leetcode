package golang

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args []struct {
		left  int
		right int
	}
	tests := []struct {
		name  string
		nums  []int
		args  args
		wants []int
	}{
		{
			name: "e1",
			nums: []int{-2, 0, 3, -5, 2, -1},
			args: args{
				{
					left:  0,
					right: 2,
				},
				{
					left:  2,
					right: 5,
				},
				{
					left:  0,
					right: 5,
				},
			},
			wants: []int{1, -1, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			na := Constructor2(tt.nums)
			for i, arg := range tt.args {
				if got := na.SumRange(arg.left, arg.right); got != tt.wants[i] {
					t.Errorf("SumRange() = %v, want %v", got, tt.wants[i])
				}
			}
		})
	}
}
