package golang

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				s: "abciiidef",
				k: 3,
			},
			want: 3,
		},
		{
			name: "e2",
			args: args{
				s: "aeiou",
				k: 2,
			},
			want: 2,
		},
		{
			name: "e3",
			args: args{
				s: "leetcode",
				k: 3,
			},
			want: 2,
		},
		{
			name: "e4",
			args: args{
				s: "rhythms",
				k: 4,
			},
			want: 0,
		},
		{
			name: "e5",
			args: args{
				s: "tryhard",
				k: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
