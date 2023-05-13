package romantointeger

import (
	"strconv"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "III",
			want: 3,
		},
		{
			s:    "IV",
			want: 4,
		},
		{
			s:    "XIV",
			want: 14,
		},
		{
			s:    "",
			want: 0,
		},
		// TODO: Add test cases.
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
