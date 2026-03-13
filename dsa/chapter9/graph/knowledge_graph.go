package graph

import "fmt"

type Class struct {
	Name string
}

type KnowledgeGraph struct {
	GraphNodes map[Class]struct{}
	Links      map[Class]map[Class]struct{}
}

func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		GraphNodes: map[Class]struct{}{},
		Links:      map[Class]map[Class]struct{}{},
	}
}

func (knowledgeGraph *KnowledgeGraph) AddClass(class Class) bool {
	if _, exists := knowledgeGraph.GraphNodes[class]; exists {
		return true
	}
	knowledgeGraph.GraphNodes[class] = struct{}{}
	return true
}

func (knowledgeGraph *KnowledgeGraph) AddLink(class1 Class, class2 Class) {
	knowledgeGraph.AddClass(class1)
	knowledgeGraph.AddClass(class2)

	if _, exists := knowledgeGraph.Links[class1]; !exists {
		knowledgeGraph.Links[class1] = make(map[Class]struct{})
	}
	knowledgeGraph.Links[class1][class2] = struct{}{}
}

func (knowledgeGraph *KnowledgeGraph) PrintLinks() {
	car := Class{
		Name: "car",
	}
	fmt.Printf("Printing all links adjacent to %s\n", car.Name)
	var node Class
	for node = range knowledgeGraph.Links[car] {
		fmt.Printf("Link: %s -> %s\n", car.Name, node.Name)
	}
	var m map[Class]struct{}
	fmt.Println("Printing all links.")
	for car, m = range knowledgeGraph.Links {
		var vertex Class
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", car.Name, vertex.Name)
		}
	}
}
