package golang

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: 3,
		},
		{
			name: "e2",
			args: args{
				nums: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "e3",
			args: args{
				nums: []int{7, 8, 9, 11, 12},
			},
			want: 1,
		},
		{
			name: "e4",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 6,
		},
		{
			name: "e5",
			args: args{
				nums: []int{1},
			},
			want: 2,
		},
		{
			name: "wa1",
			args: args{
				nums: []int{0, -1, 3, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
