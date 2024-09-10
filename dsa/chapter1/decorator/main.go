package main

import "fmt"

func main() {
	veggie := &VeggieMania{}

	veggieWithCheese := &CheeseTopping{
		pizza: veggie,
	}

	veggieWithCheeseAndTomato := &TomatoTopping{
		pizza: veggieWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", veggieWithCheeseAndTomato.getPrice())
}
