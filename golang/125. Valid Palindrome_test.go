package golang

import "testing"

func Test_isPalindrome125(t *testing.T) {
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
				s: ",,A man, a plan, a canal: Panama..",
			},
			want: true,
		},
		{
			name: "e2",
			args: args{
				s: ",,..",
			},
			want: true,
		},
		{
			name: "e3",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "wa1",
			args: args{
				// 0 和 P 的 byte 相减也是 32
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome125(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome125() = %v, want %v", got, tt.want)
			}
		})
	}
}
