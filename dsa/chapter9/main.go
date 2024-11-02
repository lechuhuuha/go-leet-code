package main

import (
	"fmt"

	"github.com/lechuhuuha/go-leet-code/dsa/chapter9/graph"
	"github.com/lechuhuuha/go-leet-code/dsa/chapter9/matrix"
)

func main() {
	socialGraph := graph.NewSocialGraph(4)
	socialGraph.AddLink(0, 1, 1)
	socialGraph.AddLink(0, 2, 1)
	socialGraph.AddLink(1, 3, 1)
	socialGraph.AddLink(2, 4, 1)
	socialGraph.PrintLinks()

	socialNetwork := graph.NewSocialNetWork()
	var root graph.Name = graph.Name("Root")
	var john graph.Name = graph.Name("John Smith")
	var per graph.Name = graph.Name("Per Jambeck")
	var cynthia graph.Name = graph.Name("Cynthia Gibas")
	socialNetwork.AddEntity(root)
	socialNetwork.AddEntity(john)
	socialNetwork.AddEntity(per)
	socialNetwork.AddEntity(cynthia)
	socialNetwork.AddLink(root, john)
	socialNetwork.AddLink(root, per)
	socialNetwork.AddLink(root, cynthia)
	var mayo graph.Name = graph.Name("Mayo Smith")
	var lorrie graph.Name = graph.Name("Lorrie Jambeck")
	var ellie graph.Name = graph.Name("Ellie Vlocksen")
	socialNetwork.AddLink(john, mayo)
	socialNetwork.AddLink(john, lorrie)
	socialNetwork.AddLink(per, ellie)
	socialNetwork.PrintLinks()

	mapLayout := graph.NewMapLayout()
	var rootPlace graph.Place = graph.Place{Name: "Algeria", Latitude: 3, Longitude: 28}
	var netherlands graph.Place = graph.Place{Name: "Netherlands", Latitude: 5.75, Longitude: 52.5}
	var korea graph.Place = graph.Place{Name: "Korea", Latitude: 124.1, Longitude: -8.36}
	var tunisia graph.Place = graph.Place{Name: "Tunisia", Latitude: 9, Longitude: 34}
	mapLayout.AddPlace(rootPlace)
	mapLayout.AddPlace(netherlands)
	mapLayout.AddPlace(korea)
	mapLayout.AddPlace(tunisia)
	mapLayout.AddLink(rootPlace, netherlands)
	mapLayout.AddLink(rootPlace, korea)
	mapLayout.AddLink(rootPlace, tunisia)
	var singapore graph.Place = graph.Place{Name: "Singapore", Latitude: 103.8, Longitude: 1.36}
	var uae graph.Place = graph.Place{Name: "UAE", Latitude: 54, Longitude: 24}
	var japan graph.Place = graph.Place{Name: "Japan", Latitude: 139.75, Longitude: 35.68}
	mapLayout.AddLink(korea, singapore)
	mapLayout.AddLink(korea, japan)
	mapLayout.AddLink(netherlands, uae)
	mapLayout.PrintLinks()

	knowledgeGraph := graph.NewKnowledgeGraph()
	var car = graph.Class{Name: "Car"}
	var tyre = graph.Class{Name: "Tyre"}
	var door = graph.Class{Name: "Door"}
	var hood = graph.Class{Name: "Hood"}
	knowledgeGraph.AddClass(car)
	knowledgeGraph.AddClass(tyre)
	knowledgeGraph.AddClass(door)
	knowledgeGraph.AddClass(hood)
	knowledgeGraph.AddLink(car, tyre)
	knowledgeGraph.AddLink(car, door)
	knowledgeGraph.AddLink(car, hood)
	var tube = graph.Class{Name: "Tube"}
	var axle = graph.Class{Name: "Axle"}
	var handle = graph.Class{Name: "Handle"}
	var windowGlass = graph.Class{Name: `Window Glass`}
	knowledgeGraph.AddClass(tube)
	knowledgeGraph.AddClass(axle)
	knowledgeGraph.AddClass(handle)
	knowledgeGraph.AddClass(windowGlass)
	knowledgeGraph.AddLink(tyre, tube)
	knowledgeGraph.AddLink(tyre, axle)
	knowledgeGraph.AddLink(door, handle)
	knowledgeGraph.AddLink(door, windowGlass)
	knowledgeGraph.PrintLinks()

	sparseMatrix := matrix.NewSparseMatrix(3, 3)
	sparseMatrix.SetValue(1, 1, 2.0)
	sparseMatrix.SetValue(1, 3, 3.0)
	fmt.Println(sparseMatrix)
	fmt.Println(sparseMatrix.NumNonZero())
}
