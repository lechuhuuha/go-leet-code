package problems

import "testing"

func TestIsStringPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string is palindrome",
			s:    "",
			want: true,
		},
		{
			name: "single character is palindrome",
			s:    "a",
			want: true,
		},
		{
			name: "odd length palindrome",
			s:    "radar",
			want: true,
		},
		{
			name: "even length palindrome",
			s:    "abba",
			want: true,
		},
		{
			name: "not a palindrome",
			s:    "hello",
			want: false,
		},
		{
			name: "mismatch near center",
			s:    "abca",
			want: false,
		},
		{
			name: "unicode palindrome",
			s:    "たけやぶやけた",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isStringPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("isStringPalindrome(%q) = %t; want %t", tt.s, got, tt.want)
			}
		})
	}
}
