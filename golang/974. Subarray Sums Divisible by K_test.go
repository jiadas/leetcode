package golang

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				A: []int{4, 5, 0, -2, -3, 1},
				K: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
