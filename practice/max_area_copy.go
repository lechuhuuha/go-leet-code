package main

import (
	"fmt"
	"unsafe"
)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func test() {
	s := make([]int, 0, 2)
	sh := (*SliceHeader)(unsafe.Pointer(&s)) // you must define a struct of the header for this to work
	fmt.Println(sh)
	doSomething(&s)
	fmt.Println(s)
}

func doSomething(a *[]int) {
	*a = append(*a, 1)
	sh := (*SliceHeader)(unsafe.Pointer(a)) // Dereference the pointer to get the actual slice
	fmt.Println(sh)
	fmt.Println(a)
}
