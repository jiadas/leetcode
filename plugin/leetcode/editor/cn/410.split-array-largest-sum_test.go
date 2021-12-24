package golang

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				nums: []int{7, 2, 5, 10, 8},
				m:    2,
			},
			want: 18,
		},
		{
			name: "e2",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				m:    2,
			},
			want: 9,
		},
		{
			name: "e3",
			args: args{
				nums: []int{1, 4, 4},
				m:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
