package interview

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example case",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "already ordered with zeroes at end",
			nums: []int{1, 2, 0, 0},
			want: []int{1, 2, 0, 0},
		},
		{
			name: "all zeroes",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "no zeroes",
			nums: []int{4, 5, 6},
			want: []int{4, 5, 6},
		},
		{
			name: "empty input",
			nums: []int{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			got := moveZeroes(nums)
			if reflect.DeepEqual(got, tt.want) == false {
				t.Errorf("moveZeroes(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
