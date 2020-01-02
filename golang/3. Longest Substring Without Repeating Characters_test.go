package golang

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "e2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "e3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "add d",
			args: args{
				s: "abcbadcbb",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}

			if got := lengthOfLongestSubstringFD(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringFD() = %v, want %v", got, tt.want)
			}
		})
	}
}
