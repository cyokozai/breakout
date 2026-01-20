package breakout

import (
 	"github.com/sago35/koebiten"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/drivers/pixel"
)


type Game struct {
	// ticker *time.Ticker
	// timer int
	// score int
	ball_x, ball_y int
	ball_vx, ball_vy int
	paddle_x, paddle_y int
	blocks [][2]int
	state string
}

var (
	screenWidth  = 128
	screenHeight = 64
	font				 = &freemono.Bold9pt7b
	area_X       = 4
	area_Y       = 1
	areaWidth    = screenWidth - area_X * 2
	areaHeight   = screenHeight - area_Y * 2
	paddleWidth  = 2
	paddleHeight = 16
	initPaddle_X = 8
	initPaddle_Y = (areaHeight - paddleHeight) / 2
	initBall_X   = areaWidth / 2
	initBall_Y   = areaHeight / 2
	ballRadius   = 2
	ballVelocity = 2
	initBlocks_X = 16
	initBlocks_Y = 8
	blockWidth   = 8
	blockHeight  = 12
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

// var (
// 	textTimer_X int16 = 32
// 	textTimer_Y int16 = 1
// 	initTimer    			= 300
// 	textScore_X int16 = 1
// 	textScore_Y int16 = 1
// 	initScore    			= 0
// 	maxScore     			= 999
// )

func (g *Game) Reset() {
	g.ball_x   = initBall_X
	g.ball_y   = initBall_Y
	g.ball_vx  = ballVelocity
	g.ball_vy  = ballVelocity
	g.paddle_x = initPaddle_X
	g.paddle_y = initPaddle_Y
	// g.timer    = initTimer
	// g.score    = initScore
	g.state    = "playing"
}

func (g *Game) IsGameOver() bool {
	return g.state == "gameover"
}

func (g *Game) PaddleControl() {
	if koebiten.IsKeyPressed(koebiten.KeyArrowUp) {
		if g.paddle_y <= area_Y {
			g.paddle_y = area_Y
		} else {
			g.paddle_y--
		}
	}
	if koebiten.IsKeyPressed(koebiten.KeyArrowDown) {
		if g.paddle_y + paddleHeight >= areaHeight {
			g.paddle_y = areaHeight - paddleHeight
		} else {
			g.paddle_y++
		}
	}
}

func (g *Game) BallMove() {
	g.ball_x += g.ball_vx
	g.ball_y += g.ball_vy

	if g.ball_y <= area_Y || g.ball_y >= areaHeight {
		g.ball_vy = -g.ball_vy
	}
	if g.ball_x >= areaWidth {
		g.ball_vx = -g.ball_vx
	}
	
	if g.ball_y + ballRadius >= g.paddle_y && g.ball_y - ballRadius <= g.paddle_y + paddleHeight {
		if g.ball_x + ballRadius > g.paddle_x + paddleWidth && g.ball_x - ballRadius < g.paddle_x + paddleWidth {
			if g.ball_y > 0 {
				g.ball_vx = -g.ball_vx
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
	koebiten.DrawRect(screen, area_X, area_Y, areaWidth, areaHeight, white)
	koebiten.DrawFilledRect(screen, g.paddle_x, g.paddle_y, paddleWidth, paddleHeight, white)
	koebiten.DrawFilledCircle(screen, g.ball_x, g.ball_y, ballRadius, white)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	game := &Game{
		// ticker: time.NewTicker(time.Second),
		// timer: initTimer,
		// score: initScore,
		ball_x: initBall_X,
		ball_y: initBall_Y,
		ball_vx: ballVelocity,
		ball_vy: ballVelocity,
		paddle_x: initPaddle_X,
		paddle_y: initPaddle_Y,
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
