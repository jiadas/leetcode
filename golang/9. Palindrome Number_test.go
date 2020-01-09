package golang

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "e2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "e3",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "number of digits is odd",
			args: args{
				x: 1221,
			},
			want: true,
		},
		{
			name: "int32 overflow",
			args: args{
				x: 2147483647,
			},
			want: false,
		},
		{
			name: "int64 overflow",
			args: args{
				x: 9223372036854775739,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
