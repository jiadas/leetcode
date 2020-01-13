package golang

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "e1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "e2",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "e3",
			args: args{
				num: 9,
			},
			want: "IX",
		},
		{
			name: "e4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "e4",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "10",
			args: args{
				num: 10,
			},
			want: "X",
		},
		{
			name: "20",
			args: args{
				num: 20,
			},
			want: "XX",
		},
		{
			name: "600",
			args: args{
				num: 600,
			},
			want: "DC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}

			if got := intToRomanFD(tt.args.num); got != tt.want {
				t.Errorf("intToRomanFD() = %v, want %v", got, tt.want)
			}
		})
	}
}
