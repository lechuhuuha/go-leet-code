package main

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color + " painted"
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{
		color: "green",
	}
}
