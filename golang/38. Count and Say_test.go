package golang

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "e1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "e2",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "e3",
			args: args{
				n: 3,
			},
			want: "21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
