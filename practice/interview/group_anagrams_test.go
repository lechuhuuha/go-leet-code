package interview

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testGroupAnagramFunc(t, "groupAnagrams", groupAnagrams)
}

func TestGroupAnagramsOptimized(t *testing.T) {
	testGroupAnagramFunc(t, "groupAnagramsOptimized", groupAnagramsOptimized)
}

func testGroupAnagramFunc(t *testing.T, name string, fn func([]string) [][]string) {
	t.Helper()

	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example case",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name: "empty input",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "single word",
			strs: []string{"abc"},
			want: [][]string{{"abc"}},
		},
		{
			name: "duplicate empty strings stay grouped",
			strs: []string{"", ""},
			want: [][]string{{"", ""}},
		},
	}

	for _, tt := range tests {
		t.Run(name+"/"+tt.name, func(t *testing.T) {
			got := fn(tt.strs)
			if reflect.DeepEqual(normalizeAnagramGroups(got), normalizeAnagramGroups(tt.want)) == false {
				t.Errorf("%s(%v) = %v; want %v", name, tt.strs, got, tt.want)
			}
		})
	}
}

func normalizeAnagramGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i, group := range groups {
		cloned := append([]string(nil), group...)
		sort.Strings(cloned)
		normalized[i] = cloned
	}

	sort.Slice(normalized, func(i, j int) bool {
		return strings.Join(normalized[i], ",") < strings.Join(normalized[j], ",")
	})

	return normalized
}
