package breakout

import (
	"time"
	// "math/rand/v2"

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

func (g *Game) Draw(screen *koebiten.Image) {
	koebiten.Println("Hello world", g.count)
}

func NewGame() *Game {
	game := &Game{}
	
	return game
}
