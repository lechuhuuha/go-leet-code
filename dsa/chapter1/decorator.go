package main

import "fmt"

type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("Process class process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}
