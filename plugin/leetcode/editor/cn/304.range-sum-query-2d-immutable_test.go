package golang

import "testing"

func TestNumMatrix_SumRegion(t *testing.T) {
	type arg struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name   string
		matrix [][]int
		args   []arg
		wants  []int
	}{
		{
			name: "e1",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			args: []arg{
				{
					row1: 2,
					col1: 1,
					row2: 4,
					col2: 3,
				},
				{
					row1: 1,
					col1: 1,
					row2: 2,
					col2: 2,
				},
				{
					row1: 1,
					col1: 2,
					row2: 2,
					col2: 4,
				},
			},
			wants: []int{8, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := Constructor3(tt.matrix)
			for i, a := range tt.args {
				if got := nm.SumRegion(a.row1, a.col1, a.row2, a.col2); got != tt.wants[i] {
					t.Errorf("SumRegion() = %v, want %v", got, tt.wants[i])
				}
			}
		})
	}
}
