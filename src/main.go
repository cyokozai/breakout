package main

import (
	"log"

	"github.com/sago35/koebiten"
	"github.com/sago35/koebiten/games/blocks/blocks"
	"github.com/sago35/koebiten/hardware"

	"github.com/cyokozai/breakout/src/breakout/breakout"
)


func main() {
	koebiten.SetHardware(hardware.Device)
	// koebiten.SetRotation(koebiten.Rotation90)
	// koebiten.SetWindowSize(64, 128)
	// koebiten.SetWindowTitle("Breakout in Go")

	game := breakout.NewGame()

	if err := koebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
