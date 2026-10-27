package main

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Result[T any] struct {
	Value T
	Err   error
}
