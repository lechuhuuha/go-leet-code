package main

import (
	"fmt"
	"sort"
)

func main() {
	linkedList := CreateLinkedList()
	ReverseLinkedList(linkedList)

	circularQueue := NewQueue(5)
	circularQueue.Add(1)
	circularQueue.Add(2)
	circularQueue.Add(3)
	circularQueue.Add(4)
	circularQueue.Add(5)

	var employees = []Employee{
		{"Graham", "231", 235643, 31},
		{"John", "3434", 245643, 42},
		{"Michael", "8934", 32432, 17},
		{"Jenny", "24334", 32444, 26},
	}
	fmt.Println(employees)
	sort.Sort(SortByAge(employees))
	fmt.Println(employees)
	sort.Slice(employees, func(i int, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	fmt.Println(employees)

	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}

	name := func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.name < Thing2.name
	}

	mass := func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.mass < Thing2.mass
	}
	distance := func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.distance < Thing2.distance
	}
	var decreasingDistance func(*Thing, *Thing) bool
	decreasingDistance = func(p1, p2 *Thing) bool {
		return distance(p2, p1)
	}
	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)
	ByFactor(mass).Sort(Things)
	fmt.Println("By mass:", Things)
	ByFactor(distance).Sort(Things)
	fmt.Println("By distance:", Things)
	ByFactor(decreasingDistance).Sort(Things)
	fmt.Println("By decreasing distance:", Things)
}