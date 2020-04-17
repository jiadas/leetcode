package golang

import "testing"

func Test_removeDuplicateStars(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				p: "a****bc***cc",
			},
			want: "a*bc*cc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateStars(tt.args.p); got != tt.want {
				t.Errorf("removeDuplicateStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch44(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "e2",
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
		},
		{
			name: "e3",
			args: args{
				s: "cb",
				p: "?a",
			},
			want: false,
		},
		{
			name: "e4",
			args: args{
				s: "adceb",
				p: "*a*b",
			},
			want: true,
		},
		{
			name: "e5",
			args: args{
				s: "acdcb",
				p: "a*c?b",
			},
			want: false,
		},
		{
			name: "e6",
			args: args{
				s: "acdcb",
				p: "*",
			},
			want: true,
		},
		{
			name: "wa1",
			args: args{
				s: "ho",
				p: "**ho",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch44(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch44() = %v, want %v", got, tt.want)
			}
		})
	}
}
