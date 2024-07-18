package main

import (
	"fmt"
	"sync"
	"time"
)

// func init() {
// 	fifo()
// }

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func cube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n * n
		}
	}()
	return out
}

func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(in))
	for _, c := range in {
		go output(c)
	}

	go func() {
		wg.Wait()
	}()

	return out

}
func fifo() {
	fmt.Println("Start Fan Out ")
	// Producer : Fan Out
	c1 := generator(1, 2, 3)
	c2 := generator(4, 5, 6)

	// Fan In
	out := merge(c1, c2)
	go func() {
		for n := range out {
			fmt.Println(n)
		}
	}()

	time.Sleep(time.Second * 2)
}
