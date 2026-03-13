package main

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 7
}
