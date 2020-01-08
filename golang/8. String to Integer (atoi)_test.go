package golang

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "e2",
			args: args{
				str: "   -42",
			},
			want: -42,
		},
		{
			name: "e3",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "e4",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "e5",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "empty str",
			args: args{
				str: "",
			},
			want: 0,
		},
		{
			name: "only space",
			args: args{
				str: "  ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
