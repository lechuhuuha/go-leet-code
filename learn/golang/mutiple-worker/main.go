package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	list := flag.String("list", "", "")
	flag.Parse()

	stringNums := strings.Split(*list, ",")
	nums := []int{}
	for _, s := range stringNums {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	fmt.Println(*list)
	total := 0
	workerInstance := 2
	totalItems := len(nums) // 6

	// accept a list of int
	// create 2 worker chan
	// divide the list to each worker
	// each worker will receive and sum the value and push to chanResult
	// main routine will wait and sum the final
	chanResult := make(chan int, len(nums))
	chanWait := make(chan int, workerInstance)
	chanErr := make(chan string, len(nums))
	chanErrDone := make(chan struct{})
	// in order to dynamic load the list into worker
	// divide the array len with the total workerInstance
	// 6 / 2 = 3
	// so we will split the array into 2 based on worker instance

	for i := 0; i < workerInstance; i++ {
		start := (i * totalItems) / workerInstance
		end := ((i + 1) * totalItems) / workerInstance
		list := nums[start:end]

		if len(list) == 0 {
			continue
		}

		go func(list []int) {
			worker(list, chanResult, chanErr)
			chanWait <- 1 // wg.add(1)
		}(list)
	}

	go func() {
		for i := 0; i < workerInstance; i++ {
			<-chanWait // wg.done()
		}
		close(chanWait)
		close(chanErr)
		close(chanResult)

	}()

	go func() {
		for err := range chanErr {
			fmt.Println(err)
		}
		close(chanErrDone)
	}()

	for v := range chanResult {
		total += v
	}

	<-chanErrDone

	fmt.Println("total: ", total)
}

func worker(nums []int, chanResult chan int, chanErr chan string) {
	result := 0
	for _, n := range nums {
		result += n
		if n == 3 {
			chanErr <- "error: this number is 3"
		}
		if n == 4 {
			chanErr <- "error: this number is 4"
		}
	}

	chanResult <- result
}
