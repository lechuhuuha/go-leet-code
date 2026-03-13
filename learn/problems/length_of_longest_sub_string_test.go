package problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
			name: "all unique characters",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "example repeated at start",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "example with duplicate inside window",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "all same characters",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "repeat forces left pointer jump",
			s:    "abba",
			want: 2,
		},
		{
			name: "longest window appears near end",
			s:    "dvdf",
			want: 3,
		},
		{
			name: "repeat after window has moved",
			s:    "tmmzuxt",
			want: 5,
		},
		{
			name: "unicode runes",
			s:    "我爱编程我",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}
