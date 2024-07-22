package systems

import (
	"errors"
	"go-tetris/entities"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	GAME_SPEED = 1
)

type GamePlaySystem struct {
	game       *entities.Game
	eventQueue chan termbox.Event
	shouldQuit bool
}

func (gs *GamePlaySystem) Init(g *entities.Game) {
	gs.game = g
	gs.eventQueue = make(chan termbox.Event)
	go func() {
		for {
			gs.eventQueue <- termbox.PollEvent()
		}
	}()
}
func (gs *GamePlaySystem) Tick(dt time.Duration) error {
	if gs.shouldQuit {
		return errors.New("quit game")
	}
	gs.play(gs.game, dt)
	return nil
}
func (gs *GamePlaySystem) Close() {

}
func (gs *GamePlaySystem) play(g *entities.Game, dt time.Duration) {
	select {
	case ev := <-gs.eventQueue:
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowLeft:
			case ev.Key == termbox.KeyArrowRight:
			case ev.Key == termbox.KeyArrowUp:
			case ev.Key == termbox.KeyArrowDown:
			case ev.Key == termbox.KeySpace:
			case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
				gs.shouldQuit = true
			}
		}
	default:
		// no event
	}
}

func NewGamePlaySystem() GamePlaySystem {
	return GamePlaySystem{}
}
