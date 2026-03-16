package problems

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string is valid",
			s:    "",
			want: true,
		},
		{
			name: "single pair parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "multiple bracket types in sequence",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "properly nested brackets",
			s:    "{[]}",
			want: true,
		},
		{
			name: "deep nesting stays valid",
			s:    "([{}])",
			want: true,
		},
		{
			name: "closing bracket before opening is invalid",
			s:    "]",
			want: false,
		},
		{
			name: "mismatched bracket types",
			s:    "(]",
			want: false,
		},
		{
			name: "wrong nesting order",
			s:    "([)]",
			want: false,
		},
		{
			name: "missing closing bracket",
			s:    "(((",
			want: false,
		},
		{
			name: "contains unsupported character",
			s:    "(a)",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %t; want %t", tt.s, got, tt.want)
			}
		})
	}
}
