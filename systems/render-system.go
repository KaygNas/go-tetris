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

	// render the board
	boardContainer := g.Board.Container
	for _, b := range boardContainer.Children {
		bb := boardContainer.GetChildAbsoluteBoundingBox(&b)
		rasterizeRect(bb.MinX, bb.MinY, bb.Width, bb.Height, b.Color)
	}

	pieceContainer := g.Piece.Container
	for _, b := range pieceContainer.Children {
		bb := pieceContainer.GetChildAbsoluteBoundingBox(&b)
		rasterizeRect(bb.MinX, bb.MinY, bb.Width, bb.Height, b.Color)
	}
	termbox.Flush()
}

// TODO: Implement the rasterizeRect function
func rasterizeRect(x, y, w, h float64, color components.Color) {
	// rasterize the block
	colorAttr := termbox.RGBToAttribute(color.R, color.G, color.B)
	// 2x because each block is 2 characters high so that it looks like a square
	for i := 0; i < int(w); i++ {
		for j := 0; j < int(h); j++ {
			cx := (int(x) + i) * 2
			cy := int(y) + j
			termbox.SetCell(cx, cy, ' ', colorAttr, colorAttr)
			termbox.SetCell(cx+1, cy, ' ', colorAttr, colorAttr)
		}
	}
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}
