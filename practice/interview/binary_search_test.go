package interview

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "finds target in middle",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "finds first element",
			nums:   []int{1, 2, 3, 4},
			target: 1,
			want:   0,
		},
		{
			name:   "finds last element",
			nums:   []int{1, 2, 3, 4},
			target: 4,
			want:   3,
		},
		{
			name:   "returns minus one when target missing",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "empty slice returns minus one",
			nums:   []int{},
			target: 7,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int{8},
			target: 8,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("binarySearch(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
