package graph

import (
	"testing"
)

func TestAddLink(t *testing.T) {
	tests := []struct {
		name          string
		graphSize     int
		linksToAdd    []Link
		expectedLinks [][]Link
	}{
		{
			name:      "Single Link",
			graphSize: 3,
			linksToAdd: []Link{
				{Vertex1: 0, Vertex2: 1, LightWeight: 5},
			},
			expectedLinks: [][]Link{
				{{Vertex1: 0, Vertex2: 1, LightWeight: 5}},
				{},
				{},
			},
		},
		{
			name:      "Multiple Links",
			graphSize: 3,
			linksToAdd: []Link{
				{Vertex1: 0, Vertex2: 1, LightWeight: 5},
				{Vertex1: 1, Vertex2: 2, LightWeight: 3},
				{Vertex1: 0, Vertex2: 2, LightWeight: 8},
			},
			expectedLinks: [][]Link{
				{{Vertex1: 0, Vertex2: 1, LightWeight: 5}, {Vertex1: 0, Vertex2: 2, LightWeight: 8}},
				{{Vertex1: 1, Vertex2: 2, LightWeight: 3}},
				{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewSocialGraph(tt.graphSize)
			for _, link := range tt.linksToAdd {
				graph.AddLink(link.Vertex1, link.Vertex2, link.LightWeight)
			}
			for i, links := range graph.Links {
				if len(links) != len(tt.expectedLinks[i]) {
					t.Errorf("vertex %d: expected %d links, got %d", i, len(tt.expectedLinks[i]), len(links))
					continue
				}
				for j, link := range links {
					if link != tt.expectedLinks[i][j] {
						t.Errorf("vertex %d: expected link %v, got %v", i, tt.expectedLinks[i][j], link)
					}
				}
			}
		})
	}
}
