package problems

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "finds target range in middle",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "returns not found when target absent",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "returns not found for empty input",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			name:   "single element match",
			nums:   []int{8},
			target: 8,
			want:   []int{0, 0},
		},
		{
			name:   "two elements both match",
			nums:   []int{8, 8},
			target: 8,
			want:   []int{0, 1},
		},
		{
			name:   "target spans duplicate block",
			nums:   []int{1, 2, 2, 2, 3},
			target: 2,
			want:   []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
