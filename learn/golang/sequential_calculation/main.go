package main

import "fmt"

func main() {
	numbers := []int64{
		1, 2, 3, 4, 5, 5, 6, 7, 8, 9,
	}
	result := sumConcurrent(numbers, 8)
	fmt.Println(result)
}

func sum(numbers []int64) int64 {
	var result int64
	for _, n := range numbers {
		result += n
	}
	return int64(result)
}

func sumConcurrent(numbers []int64, workerCount int) int64 {
	resultChan := make(chan int64, workerCount)
	done := make(chan bool)

	mid := len(numbers) / 2
	firstHalf := numbers[:mid]
	secondHalf := numbers[mid:]

	var result int64

	go func() {
		sumSplit(firstHalf, resultChan)
		done <- true

	}()
	go func() {
		sumSplit(secondHalf, resultChan)
		done <- true

	}()

	<-done
	<-done
	close(resultChan)
	for r := range resultChan {
		result += r
	}
	return result
}

func sumSplit(numbers []int64, resultChan chan<- int64) {
	var result int64
	for _, n := range numbers {
		result += n
	}
	resultChan <- result
}
