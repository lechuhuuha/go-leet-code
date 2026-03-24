package problems

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "example case",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "empty input",
			temperatures: []int{},
			want:         []int{},
		},
		{
			name:         "single day has no warmer future day",
			temperatures: []int{80},
			want:         []int{0},
		},
		{
			name:         "strictly increasing temperatures",
			temperatures: []int{70, 71, 72, 73},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "strictly decreasing temperatures",
			temperatures: []int{73, 72, 71, 70},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "same temperatures do not count as warmer",
			temperatures: []int{70, 70, 70},
			want:         []int{0, 0, 0},
		},
		{
			name:         "warmer day appears after repeated lower temperatures",
			temperatures: []int{70, 71, 71, 72},
			want:         []int{1, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures(%v) = %v; want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
