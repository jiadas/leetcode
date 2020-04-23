package golang

import "testing"

func Test_findMin154(t *testing.T) {
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
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "e2",
			args: args{
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "e3",
			args: args{
				nums: []int{2, 2, 2, 0, 2, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin154(tt.args.nums); got != tt.want {
				t.Errorf("findMin154() = %v, want %v", got, tt.want)
			}
		})
	}
}
