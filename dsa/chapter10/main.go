package main

import (
	"fmt"

	"github.com/lechuhuuha/go-leet-code/dsa/chapter10/cachemanagement"
	"github.com/lechuhuuha/go-leet-code/dsa/chapter10/garbagecollection"
	"github.com/lechuhuuha/go-leet-code/dsa/chapter10/referencecounting"
	"github.com/lechuhuuha/go-leet-code/dsa/chapter10/spaceallocation"
)

func main() {
	var stack *garbagecollection.Stack = &garbagecollection.Stack{}
	stack.New()
	var reference1 = garbagecollection.NewReferenceCounter()
	var reference2 = garbagecollection.NewReferenceCounter()
	var reference3 = garbagecollection.NewReferenceCounter()
	var reference4 = garbagecollection.NewReferenceCounter()
	stack.Push(reference1)
	stack.Push(reference2)
	stack.Push(reference3)
	stack.Push(reference4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())

	referenceCounter := referencecounting.NewReferenceCounter()
	referenceCounter.Add()
	fmt.Println(referenceCounter.Num)

	cacheMangager := cachemanagement.NewCache()
	cacheMangager.Set("name", "John Smith", 2000000)

	name := cacheMangager.Get("name")
	fmt.Println(name)

	number := 17
	fmt.Println("value of number", number, "Address of number",
		&number)
	spaceallocation.AddOne(number)
	fmt.Println("value of number after adding One", number, "Address of", &number)
}
