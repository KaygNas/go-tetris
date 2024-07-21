package systems

import "go-tetris/entities"

type IGamePlaySystem interface {
	System
	play(*entities.Game)
}

type GamePlaySystem struct {
}

func (gs *GamePlaySystem) Init(g *entities.Game) {

}
func (gs *GamePlaySystem) Tick() {

}
func (gs *GamePlaySystem) Close() {

}
func (gs *GamePlaySystem) play(g *entities.Game) {

}

func NewGamePlaySystem() IGamePlaySystem {
	return &GamePlaySystem{}
}
