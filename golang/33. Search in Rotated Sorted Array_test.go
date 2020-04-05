package golang

import "testing"

func Test_search33(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "e2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "wa1",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "wa2",
			args: args{
				nums:   []int{1, 3},
				target: 1,
			},
			want: 0,
		},
		{
			name: "wa3",
			args: args{
				nums:   []int{1, 3},
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search33(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search33() = %v, want %v", got, tt.want)
			}
		})
	}
}
