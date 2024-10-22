package main

import "fmt"

type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

func (e *Employee) String() string {
	return fmt.Sprintf("%s: %d,%s,%d", e.Name,
		e.Age, e.ID,
		e.SSN)
}

type SortByAge []Employee

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
