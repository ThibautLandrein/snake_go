package snake

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  int = 600
	ScreenHeight int = 600
	boardRows    int = 15
	boardCols    int = 15
)

var (
	backgroundColor = color.RGBA{50, 100, 50, 50}
	snakeColor      = color.RGBA{200, 50, 150, 150}
	foodColor       = color.RGBA{200, 200, 50, 150}
	headColor       = color.RGBA{0, 0, 255, 150}
)

type Game struct {
	input *Input
	board *Board
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Update() error {
	return g.board.Update(g.input)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	if g.board.gameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Over. Score: %d", g.board.points))
	} else {
		width := ScreenHeight / boardRows

		p := g.board.snake.body[len(g.board.snake.body)-1]
		ebitenutil.DrawRect(screen, float64(p.y*width), float64(p.x*width), float64(width), float64(width), headColor)
		for _, p := range g.board.snake.body[1:] {
			ebitenutil.DrawRect(screen, float64(p.y*width), float64(p.x*width), float64(width), float64(width), snakeColor)
		}
		if g.board.food != nil {
			ebitenutil.DrawRect(screen, float64(g.board.food.y*width), float64(g.board.food.x*width), float64(width), float64(width), foodColor)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.points))
	}
}
