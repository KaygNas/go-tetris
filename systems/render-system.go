package systems

import (
	"go-tetris/entities"
	"time"

	"github.com/nsf/termbox-go"
)

type RenderSystem struct {
	game *entities.Game
}

func (rs *RenderSystem) Init(g *entities.Game) {
	termbox.Init()
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
		termbox.SetCell(x, y, ' ', termbox.Attribute(b.Color), termbox.Attribute(b.Color))
	}
	termbox.Flush()
}

func NewRenderSystem() RenderSystem {
	return RenderSystem{}
}
