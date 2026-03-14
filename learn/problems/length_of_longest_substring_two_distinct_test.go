package problems

import "testing"

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "two distinct characters entire string",
			s:    "ece",
			want: 3,
		},
		{
			name: "example requires shrinking window",
			s:    "eceba",
			want: 3,
		},
		{
			name: "long run of two characters",
			s:    "ccaabbb",
			want: 5,
		},
		{
			name: "all same characters",
			s:    "aaaaa",
			want: 5,
		},
		{
			name: "all unique characters",
			s:    "abcde",
			want: 2,
		},
		{
			name: "window grows again after third character removed",
			s:    "abaccc",
			want: 4,
		},
		{
			name: "unicode runes",
			s:    "我我爱爱编",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstringTwoDistinct(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}
