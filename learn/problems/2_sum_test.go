package problems

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "happy path",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "already first two elements",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "duplicate values can form answer",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "includes negative numbers",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "includes zeros and duplicate zero pair",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "harder case match appears late",
			nums:   []int{1, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, -7, 4},
			target: -3,
			want:   []int{12, 13},
		},
		{
			name:   "multiple possible pairs returns first discovered by scan",
			nums:   []int{1, 5, 1, 5},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "no solution returns empty slice",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
