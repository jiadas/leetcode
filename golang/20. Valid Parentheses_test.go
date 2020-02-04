package golang

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "e2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "e3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "e4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "e3",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
