package w186

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				s: "011101",
			},
			want: 5,
		},
		{
			name: "e2",
			args: args{
				s: "00111",
			},
			want: 5,
		},
		{
			name: "e3",
			args: args{
				s: "1111",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
