package main

import "sort"

type Mass float64
type Miles float64

type Thing struct {
	name          string
	mass          Mass
	distance      Miles
	meltingPoint  int
	freezingPoint int
}

type ByFactor func(Thing1 *Thing, Thing2 *Thing) bool

type ThingSorter struct {
	Things   []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool
}

// Len implements sort.Interface.
func (t ThingSorter) Len() int {
	return len(t.Things)
}

// Less implements sort.Interface.
func (t ThingSorter) Less(i int, j int) bool {
	return t.byFactor(&t.Things[i],
		&t.Things[j])
}

// Swap implements sort.Interface.
func (t ThingSorter) Swap(i int, j int) {
	t.Things[i], t.Things[j] = t.Things[j],
		t.Things[i]
}

func (b ByFactor) Sort(things []Thing) {
	sortedThings := ThingSorter{
		Things:   things,
		byFactor: b,
	}
	sort.Sort(sortedThings)
}
