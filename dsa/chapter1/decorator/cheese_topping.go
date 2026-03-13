package main

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	return c.pizza.getPrice() + 10
}
