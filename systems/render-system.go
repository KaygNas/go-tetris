package systems

import (
	"go-tetris/entities"
	"time"

	"github.com/nsf/termbox-go"
)

type IRenderSystem interface {
	System
	render(*entities.Game)
}

type RenderSystem struct {
	game *entities.Game
}

func (rs *RenderSystem) Init(g *entities.Game) {
	termbox.Init()
	rs.game = g
}
func (rs *RenderSystem) Tick() {
	rs.render(rs.game)
}
func (rs *RenderSystem) Close() {
	termbox.Close()
}
func (rs *RenderSystem) render(g *entities.Game) {
	now := time.Now().String()
	for i, c := range now {
		termbox.SetCell(0+i, 0, c, termbox.ColorWhite, termbox.ColorBlack)
	}
	termbox.Flush()
}

func NewRenderSystem() IRenderSystem {
	return &RenderSystem{}
}
