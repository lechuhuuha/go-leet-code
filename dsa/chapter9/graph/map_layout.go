package graph

import "fmt"

type (
	Place struct {
		Name      string
		Latitude  float64
		Longitude float64
	}

	MapLayout struct {
		GraphNodes map[Place]struct{}
		Links      map[Place]map[Place]struct{}
	}
)

func NewMapLayout() *MapLayout {
	return &MapLayout{
		GraphNodes: map[Place]struct{}{},
		Links:      map[Place]map[Place]struct{}{},
	}
}

func (m *MapLayout) AddPlace(p Place) bool {
	if _, exists := m.GraphNodes[p]; exists {
		return true
	}

	m.GraphNodes[p] = struct{}{}
	return true
}

func (m *MapLayout) AddLink(place1, place2 Place) {
	m.AddPlace(place1)
	m.AddPlace(place2)

	if _, exists := m.Links[place1]; !exists {
		m.Links[place1] = make(map[Place]struct{})
	}
	m.Links[place1][place2] = struct{}{}
}

func (m *MapLayout) PrintLinks() {
	root := Place{"Algeria", 3, 28}

	fmt.Printf("Printing all links adjacent to %s\n", root.Name)

	for node := range m.Links[root] {
		fmt.Printf("Link: %s -> %s\n", root.Name, node.Name)
	}

	fmt.Println("Printing all links.")

	for root, m := range m.Links {
		var vertex Place
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", root.Name, vertex.Name)
		}
	}
}
