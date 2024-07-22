package systems

import (
	"errors"
	"go-tetris/entities"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	FALLING_SPEED = 1 * time.Second
)

type GamePlaySystem struct {
	game         *entities.Game
	eventQueue   chan termbox.Event
	fallingTimer *time.Timer
	shouldQuit   bool
}

func (gs *GamePlaySystem) Init(g *entities.Game) {
	gs.game = g

	gs.eventQueue = make(chan termbox.Event)
	go func() {
		for {
			gs.eventQueue <- termbox.PollEvent()
		}
	}()

	gs.fallingTimer = time.NewTimer(FALLING_SPEED)
}
func (gs *GamePlaySystem) Tick(dt time.Duration) error {
	if gs.shouldQuit {
		return errors.New("quit game")
	}
	gs.play(gs.game, dt)
	return nil
}
func (gs *GamePlaySystem) Close() {
	termbox.Close()
	gs.fallingTimer.Stop()
}
func (gs *GamePlaySystem) play(g *entities.Game, dt time.Duration) {
	select {
	case ev := <-gs.eventQueue:
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyArrowLeft:
				g.Piece.MoveLeft()
			case ev.Key == termbox.KeyArrowRight:
				g.Piece.MoveRight()
			case ev.Key == termbox.KeyArrowUp:
				g.Piece.RotateCW()
			case ev.Key == termbox.KeyArrowDown:
				g.Piece.MoveDown()
				gs.fallingTimer.Reset(FALLING_SPEED) // reset timer so that select below will not be triggered
			case ev.Key == termbox.KeySpace:

			case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
				gs.shouldQuit = true
			}
		}
	default:
		// no event
	}

	select {
	case <-gs.fallingTimer.C:
		g.Piece.MoveDown()
		gs.fallingTimer.Reset(FALLING_SPEED)
	default:
		// no event
	}
}

func NewGamePlaySystem() GamePlaySystem {
	return GamePlaySystem{}
}
