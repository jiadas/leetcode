package golang

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "e1",
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPowIter(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "e1",
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPowIter(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPowIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
