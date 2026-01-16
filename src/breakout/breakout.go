package breakout

import (
	"time"
	"math/rand"
	
 	"github.com/sago35/koebiten"
 	"github.com/sago35/koebiten/hardware"
)


type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	return nil
}

func NewGame(game *Game) *Game {
	rand.Seed(time.Now().UnixNano())
	game := &Game{}


	return game
}