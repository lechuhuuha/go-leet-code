package graph

import "fmt"

type Name string

type SocialNetwork struct {
	GraphNodes map[Name]struct{}
	Links      map[Name]map[Name]struct{}
}

func NewSocialNetWork() *SocialNetwork {
	return &SocialNetwork{
		GraphNodes: map[Name]struct{}{},
		Links:      map[Name]map[Name]struct{}{},
	}
}

func (s *SocialNetwork) AddEntity(name Name) bool {
	if _, exists := s.GraphNodes[name]; exists {
		return exists
	}

	s.GraphNodes[name] = struct{}{}
	return true
}

func (s *SocialNetwork) AddLink(name1, name2 Name) {
	s.AddEntity(name1)
	s.AddEntity(name2)

	if _, exists := s.Links[name1]; !exists {
		s.Links[name1] = make(map[Name]struct{})
	}
	s.Links[name1][name2] = struct{}{}
}

func (s *SocialNetwork) PrintLinks() {
	root := Name("Root")
	fmt.Printf("Printing all links adjacent to %s\n", root)
	var node Name
	for node = range s.Links[root] {
		// Edge exists from u to v.
		fmt.Printf("Link: %s -> %s\n", root, node)
	}
	var m map[Name]struct{}
	fmt.Println("Printing all links.")
	for root, m = range s.Links {
		var vertex Name
		for vertex = range m {
			// Edge exists from u to v.
			fmt.Printf("Link: %s -> %s\n", root, vertex)
		}
	}
}
