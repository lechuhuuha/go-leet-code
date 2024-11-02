package graph

import "fmt"

type (
	SocialGraph struct {
		Size  int
		Links [][]Link
	}

	Link struct {
		Vertex1     int
		Vertex2     int
		LightWeight int
	}
)

func NewSocialGraph(num int) *SocialGraph {
	return &SocialGraph{
		Size:  num,
		Links: make([][]Link, num),
	}
}

func (s *SocialGraph) AddLink(vertex1, vertex2, weight int) {
	s.Links[vertex1] = append(s.Links[vertex1], Link{
		Vertex1:     vertex1,
		Vertex2:     vertex2,
		LightWeight: weight,
	})
}

func (s *SocialGraph) PrintLinks() {
	vertex := 0
	for _, link := range s.Links[vertex] {
		fmt.Printf("Link : %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LightWeight)
	}

	for _, adjacent := range s.Links {
		for _, link := range adjacent {
			fmt.Printf("Link : %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LightWeight)
		}
	}
}
