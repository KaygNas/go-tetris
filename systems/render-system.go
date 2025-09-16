package systems

import (
	"os"
	"time"

	"github.com/KaygNas/go-tetris/components"
	"github.com/KaygNas/go-tetris/entities"

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

	renderContainer(&g.Board.Container)
	renderContainer(&g.Piece.Container)
	renderContainer(&g.LockedPieces.Container)
	renderContainer(&g.Instructions.Container)
	renderContainer(&g.Stat.Container)

	termbox.Flush()
}

func renderContainer(c *components.Container) {
	for _, b := range c.Blocks {
		bb := c.GetChildAbsoluteBoundingBox(&b)
		rasterizeRect(bb.MinX, bb.MinY, bb.Width, bb.Height, b.Color, b.Char)
	}
}

// TODO: Implement the rasterizeRect function
func rasterizeRect(x, y, w, h int, color components.Color, char *rune) {
	var DEFAULT_CHAR rune = ' '
	if os.Getenv("MODE") == "development" {
		DEFAULT_CHAR = '-'
	}

	// rasterize the block
	colorAttr := termbox.RGBToAttribute(color.R, color.G, color.B)
	// 2x because each block is 2 characters high so that it looks like a square
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			cx := (x + i) * 2
			cy := y + j
			if char == nil {
				char = &DEFAULT_CHAR
				termbox.SetCell(cx, cy, *char, colorAttr, colorAttr)
				termbox.SetCell(cx+1, cy, *char, colorAttr, colorAttr)
			} else {
				termbox.SetCell(cx, cy, *char, colorAttr, termbox.ColorDefault)
			}
		}
	}
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}
