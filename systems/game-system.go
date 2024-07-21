package systems

import (
	"go-tetris/entities"
	"time"
)

const (
	GAME_SPEED = 1
)

type GamePlaySystem struct {
	game *entities.Game
}

func (gs *GamePlaySystem) Init(g *entities.Game) {
	gs.game = g
}
func (gs *GamePlaySystem) Tick(dt time.Duration) {
	gs.play(gs.game, dt)
}
func (gs *GamePlaySystem) Close() {

}
func (gs *GamePlaySystem) play(g *entities.Game, dt time.Duration) {
}

func NewGamePlaySystem() GamePlaySystem {
	return GamePlaySystem{}
}
