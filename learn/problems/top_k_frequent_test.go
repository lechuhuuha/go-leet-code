package problems

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "happy path example",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "nil input returns empty slice",
			nums: nil,
			k:    1,
			want: []int{},
		},
		{
			name: "empty input returns empty slice",
			nums: []int{},
			k:    1,
			want: []int{},
		},
		{
			name: "zero k returns empty slice",
			nums: []int{1, 1, 2, 2, 3},
			k:    0,
			want: []int{},
		},
		{
			name: "single distinct value",
			nums: []int{7, 7, 7, 7},
			k:    1,
			want: []int{7},
		},
		{
			name: "includes negative numbers and zero",
			nums: []int{-1, -1, -1, 0, 0, 2, 2, 2, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
		{
			name: "k equals number of unique values",
			nums: []int{5, 5, 4, 4, 4, 3},
			k:    3,
			want: []int{3, 4, 5},
		},
		{
			name: "larger frequencies replace earlier smaller values",
			nums: []int{9, 9, 9, 9, 8, 8, 8, 7, 7, 6},
			k:    2,
			want: []int{8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !sameIntsIgnoringOrder(got, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func sameIntsIgnoringOrder(got, want []int) bool {
	gotCopy := append([]int(nil), got...)
	wantCopy := append([]int(nil), want...)

	sort.Ints(gotCopy)
	sort.Ints(wantCopy)

	return reflect.DeepEqual(gotCopy, wantCopy)
}
