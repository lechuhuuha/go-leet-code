package main

import (
	"reflect"
	"testing"
)

func TestDivideArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want [][]int
	}{
		{
			name: "valid groups",
			nums: []int{1, 3, 5, 7, 9, 11},
			k:    4,
			want: [][]int{{1, 3, 5}, {7, 9, 11}},
		},
		{
			name: "invalid group due to k",
			nums: []int{1, 3, 9, 10, 12, 13},
			k:    3,
			want: [][]int{},
		},
		{
			name: "all identical values",
			nums: []int{2, 2, 2, 2, 2, 2},
			k:    0,
			want: [][]int{{2, 2, 2}, {2, 2, 2}},
		},
		{
			name: "reverse order input",
			nums: []int{6, 5, 4, 3, 2, 1},
			k:    2,
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name: "exact k difference",
			nums: []int{1, 2, 4, 5, 6, 7},
			k:    3,
			want: [][]int{{1, 2, 4}, {5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideArray(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
