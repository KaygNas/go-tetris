package systems

import (
	"go-tetris/components"
	"go-tetris/entities"
	"time"

	"github.com/nsf/termbox-go"
)

type RenderSystem struct {
	game *entities.Game
}

func (rs *RenderSystem) Init(g *entities.Game) {
	termbox.Init()
	termbox.SetOutputMode(termbox.OutputRGB)
	rs.game = g
}
func (rs *RenderSystem) Tick(dt time.Duration) error {
	rs.render(rs.game)
	return nil
}
func (rs *RenderSystem) Close() {
	termbox.Close()
}
func (rs *RenderSystem) render(g *entities.Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	pieceContainer := g.Piece.Container
	for _, b := range pieceContainer.Children {
		x, y := pieceContainer.GetChildAbsPosition(&b)
		rasterizePixel(x, y, b.Color)
	}
	termbox.Flush()
}

// TODO: Implement the rasterizePixel function
func rasterizePixel(x, y float64, color components.Color) {
	// rasterize the block
	colorAttr := termbox.RGBToAttribute(color.R, color.G, color.B)
	// 2x because each block is 2 characters high so that it looks like a square
	x *= 2
	termbox.SetCell(int(x), int(y), ' ', colorAttr, colorAttr)
	termbox.SetCell(int(x)+1, int(y), ' ', colorAttr, colorAttr)
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}
