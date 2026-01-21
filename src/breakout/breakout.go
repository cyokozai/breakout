package breakout

import (
	"machine"
	"time"

	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/tone"
	"tinygo.org/x/tinyfont/freemono"
)

type Game struct {
	ballX, ballY int
	ballVX, ballVY int
	paddleX, paddleY int
	blocks [][2]int
	state string
}

var (
	screenWidth  = 64
	screenHeight = 128
	font				 = &freemono.Bold9pt7b
	areaX        = 4
	areaY        = 1
	areaWidth    = screenWidth - areaX * 2
	areaHeight   = screenHeight - areaY * 2
	paddleWidth  = 16
	paddleHeight = 2
	initPaddleX  = (areaWidth - paddleWidth) / 2 + areaX
	initPaddleY  = 120
	initBallX    = areaWidth / 2
	initBallY    = areaHeight / 2
	ballRadius   = 2
	ballVelocity = 2
	initBlocksX  = 8
	initBlocksY  = 8
	blockWidth   = 6
	blockHeight  = 4
)

var (
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

var (
	messageTextX int16 = 8
	messageTextY int16 = 48
)

var (
	buzzerPin = machine.GPIO1
	PWM 	    = machine.PWM0
	mute      = tone.Note(0)
	buzzer, _ = tone.New(PWM, buzzerPin)
)


func (g *Game) PaddleControl() {
	if koebiten.IsKeyPressed(koebiten.KeyArrowUp) {
		if g.paddleX <= areaX {
			g.paddleX = areaX
		} else {
			g.paddleX--
		}
	}
	if koebiten.IsKeyPressed(koebiten.KeyArrowDown) {
		if g.paddleX + paddleWidth >= areaX + areaWidth {
			g.paddleX = areaX + areaWidth - paddleWidth
		} else {
			g.paddleX++
		}
	}
}

func (g *Game) BallMove() {
	g.ballX += g.ballVX
	g.ballY += g.ballVY

	if g.ballX <= areaX || g.ballX >= areaWidth {
		g.ballVX = -g.ballVX
	}
	if g.ballY <= 0 {
		g.ballVY = -g.ballVY
	}

	for i, b := range g.blocks {
		if g.ballX + ballRadius >= b[0] && g.ballX - ballRadius <= b[0] + blockWidth {
			if g.ballY + ballRadius >= b[1] && g.ballY - ballRadius <= b[1] + blockHeight {
				g.ballVX = -g.ballVX
				g.blocks = append(g.blocks[:i], g.blocks[i+1:]...)

				break
			}
		}
	}

	if g.ballY + ballRadius >= g.paddleY && g.ballY + ballRadius <= g.paddleY + paddleHeight {
		if g.ballX - ballRadius >= g.paddleX && g.ballX + ballRadius <= g.paddleX + paddleWidth {
			g.ballVY = -g.ballVY
			g.ballY = g.paddleY - paddleHeight - ballRadius
		}
	}
}

func buzzerPlay(buzzer tone.Speaker, state string) {
	if state == "boot" {
		buzzer.SetNote(tone.B6)
		time.Sleep(time.Millisecond * 100)
		buzzer.SetNote(mute)
	} else if state == "gameover" {
		buzzer.SetNote(tone.C6)
		time.Sleep(time.Millisecond * 300)
		buzzer.SetNote(mute)
	} else if state == "clear" {
		buzzer.SetNote(tone.E6)
		time.Sleep(time.Millisecond * 100)
		buzzer.SetNote(tone.G6)
		time.Sleep(time.Millisecond * 100)
		buzzer.SetNote(tone.C7)
		time.Sleep(time.Millisecond * 200)
		buzzer.SetNote(mute)
	}
}

func (g *Game) Update() error {
	if g.state != "playing" {
		return nil
	} else {
		g.PaddleControl()
		g.BallMove()

		if g.ballY >= areaY + areaHeight {
			g.state = "gameover"
			buzzerPlay(buzzer, "gameover")
		}
		if len(g.blocks) == 0 {
			g.state = "clear"
			buzzerPlay(buzzer, "clear")
		}
	}

	return nil
}

func (g *Game) Draw(screen *koebiten.Image) {
	if g.state == "playing" {
		koebiten.DrawRect(screen, areaX, areaY, areaWidth, areaHeight, white)
		koebiten.DrawFilledRect(screen, g.paddleX, g.paddleY, paddleWidth, paddleHeight, white)
		koebiten.DrawFilledCircle(screen, g.ballX, g.ballY, ballRadius, white)
		for _, b := range g.blocks {
			koebiten.DrawFilledRect(screen, b[0], b[1], blockWidth, blockHeight, white)
		}
	} else if g.state == "gameover" {
		koebiten.DrawFilledRect(screen, 0, 0, screenWidth, screenHeight, black)
		koebiten.DrawText(screen, "GAME\nOVER", font, messageTextX, messageTextY, white)
	} else if g.state == "clear" {
		koebiten.DrawFilledRect(screen, 0, 0, screenWidth, screenHeight, black)
		koebiten.DrawText(screen, "CLEAR!", font, messageTextX, messageTextY, white)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	buzzerPlay(buzzer, "boot")

	game := &Game{
		ballX: initBallX,
		ballY: initBallY,
		ballVX: ballVelocity,
		ballVY: ballVelocity,
		paddleX: initPaddleX,
		paddleY: initPaddleY,
		blocks: func() [][2]int {
			b := make([][2]int, 0)
			for i := 0; i < 6; i++ {
				for j := 0; j < 6; j++ {
					b = append(b, [2]int{initBlocksX + j*9, initBlocksY + i*10})
				}
			}
			return b
		}(),
		state: "playing",
	}
	
	return game
}
