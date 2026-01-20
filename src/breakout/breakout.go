package breakout

import (
	"time"

 	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
)


type Game struct {
	ticker *time.Ticker
	timer int
	score int
	ball_x, ball_y int
	ball_vx, ball_vy int
	paddle_x, paddle_y int
	blocks [][2]int
	state string
}

var (
	screenWidth  = 64
	screenHeight = 128
	paddleWidth  = 2
	paddleHeight = 16
	ballRadius   = 2
	ballVelocity = 2
	blockWidth   = 8
	blockHeight  = 12
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

func (g *Game) PaddleControl() {
	if koebiten.IsKeyPressed(koebiten.KeyArrowUp) {
		if g.paddle_y <= 0 {
			g.paddle_y = 0
		} else {
			g.paddle_y--
		}
	}
	if koebiten.IsKeyPressed(koebiten.KeyArrowDown) {
		if g.paddle_y >= screenWidth - paddleHeight {
			g.paddle_y = screenWidth - paddleHeight
		} else {
			g.paddle_y++
		}
	}
}

func (g *Game) BallMove() {
	g.ball_x += g.ball_vx
	g.ball_y += g.ball_vy

	if g.ball_y <= 0 || g.ball_y >= screenWidth {
		g.ball_vy = -g.ball_vy
	}
	if g.ball_x >= screenHeight {
		g.ball_vx = -g.ball_vx
	}
	
	if g.ball_y + ballRadius >= g.paddle_y && g.ball_y - ballRadius <= g.paddle_y + paddleHeight {
		if g.ball_x + ballRadius > g.paddle_x && g.ball_x - ballRadius < g.paddle_x + paddleWidth {
			if g.ball_y > 0 {
				g.ball_vx *= -1
				g.ball_x = g.paddle_x + paddleWidth + ballRadius
			}
		}
	}

	// if g.ball_x <= 0 {
	// 	g.state = "gameover"
	// }
}

func (g *Game) Update() error {
	if g.state != "playing" {
		return nil
	} else {
		// go func() {
		// 	for ; g.timer >= 0; g.timer-- { <-g.ticker.C }
		// }()
		// if g.timer <= 0 { g.state = "gameover" }

		g.PaddleControl()

		g.BallMove()
	}

	return nil
}

func (g *Game) Draw(screen *koebiten.Image) {
	koebiten.DrawFilledRect(screen, g.paddle_x, g.paddle_y ,paddleWidth, paddleHeight, white)
	koebiten.DrawFilledCircle(screen, g.ball_x, g.ball_y, ballRadius, white)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenHeight, screenWidth
}

func NewGame() *Game {
	game := &Game{
		ticker: time.NewTicker(time.Second),
		timer: 300,
		score: 0,
		ball_x: screenHeight / 2,
		ball_y: screenWidth / 2,
		ball_vx: ballVelocity,
		ball_vy: ballVelocity,
		paddle_x: 8,
		paddle_y: (screenWidth - paddleHeight) / 2,
		blocks: func() [][2]int {
			b := make([][2]int, 0)
			for i := 0; i < 5; i++ {
				for j := 0; j < 8; j++ {
					b = append(b, [2]int{16 + j*12, 8 + i*6})
				}
			}
			return b
		}(),
		state: "playing",
	}
	
	return game
}
