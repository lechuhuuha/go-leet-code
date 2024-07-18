package main

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maxArea(tt.numbers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v) = %v; want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
