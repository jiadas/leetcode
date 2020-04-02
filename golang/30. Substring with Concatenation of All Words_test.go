package golang

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "e1",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{0, 9},
		},
		{
			name: "e2",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
