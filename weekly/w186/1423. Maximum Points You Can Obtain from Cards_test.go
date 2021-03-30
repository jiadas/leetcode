package w186

import "testing"

func Test_maxScore2(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
				k:          3,
			},
			want: 12,
		},
		{
			name: "e2",
			args: args{
				cardPoints: []int{2, 2, 2},
				k:          2,
			},
			want: 4,
		},
		{
			name: "e3",
			args: args{
				cardPoints: []int{9, 7, 7, 9, 7, 7, 9},
				k:          7,
			},
			want: 55,
		},
		{
			name: "e4",
			args: args{
				cardPoints: []int{1, 1000, 1},
				k:          1,
			},
			want: 1,
		},
		{
			name: "e5",
			args: args{
				cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
				k:          3,
			},
			want: 202,
		},
		{
			name: "wa1",
			args: args{
				cardPoints: []int{100, 40, 17, 9, 73, 75},
				k:          3,
			},
			want: 248,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore2(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore2() = %v, want %v", got, tt.want)
			}
		})
	}
}
