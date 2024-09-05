package main

type Game struct {
	terrorists       []*Player
	counterTerrorist []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:       make([]*Player, 1),
		counterTerrorist: make([]*Player, 1),
	}
}

func (g *Game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorist = append(g.counterTerrorist, player)
}
