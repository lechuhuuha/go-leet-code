package problems

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "example case one",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			name:  "example case two",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		{
			name:  "empty nums1 returns empty output",
			nums1: []int{},
			nums2: []int{1, 2, 3},
			want:  []int{},
		},
		{
			name:  "single element without greater value",
			nums1: []int{1},
			nums2: []int{1},
			want:  []int{-1},
		},
		{
			name:  "next greater appears later in nums2",
			nums1: []int{2, 1},
			nums2: []int{2, 3, 1},
			want:  []int{3, -1},
		},
		{
			name:  "all values have a greater element except last",
			nums1: []int{1, 3, 5},
			nums2: []int{1, 2, 3, 4, 5, 6},
			want:  []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement(%v, %v) = %v; want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
