package golang

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			},
			want: 11,
		},
		{
			name: "f1",
			args: args{
				triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal1(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
