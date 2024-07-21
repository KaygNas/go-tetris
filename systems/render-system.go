package systems

import (
	"go-tetris/entities"
)

type IRenderSystem interface {
	System
	render(*entities.Game)
}

type RenderSystem struct{}

func (rs *RenderSystem) Init(g *entities.Game) {

}
func (rs *RenderSystem) Tick() {

}
func (rs *RenderSystem) Close() {

}
func (rs *RenderSystem) render(g *entities.Game) {

}

func NewRenderSystem() IRenderSystem {
	return &RenderSystem{}
}
