package systems

import (
	"errors"
	"go-tetris/entities"
	"math"
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
	g.NextPiece()
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
			case ev.Ch == 'n':
				gs.game.NewGame()
				gs.fallingTimer.Reset(FALLING_SPEED) // reset timer so that select below will not be triggered
			case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
				gs.shouldQuit = true
			}
		}
	default:
		// no event
	}

	select {
	case <-gs.fallingTimer.C:
		gs.tryLockCurrentPiece()
		gs.clearLines()
		g.Piece.MoveDown()
		gs.fallingTimer.Reset(FALLING_SPEED)
	default:
		// no event
	}

	// check if the piece is collided with the locked pieces
	if isChildrenCollided := g.LockedPieces.Container.IsChildrenCollide(&g.Piece.Container); isChildrenCollided {
		g.Piece.RestoreTransform()
	}

	// check if the piece is out of bounds
	if isOutOfBounds := !g.Board.Container.BoundingBoxContain(&g.Piece.Container); isOutOfBounds {
		g.Piece.MoveInto(&g.Board.Container)
	}

}

func (gs *GamePlaySystem) tryLockCurrentPiece() {
	g := gs.game
	// check if the piece can move down
	g.Piece.MoveDown()
	if isChildrenCollided := g.LockedPieces.Container.IsChildrenCollide(&g.Piece.Container); isChildrenCollided {
		g.Piece.RestoreTransform()
		gs.lockCurrentPiece()
	} else {
		g.Piece.RestoreTransform()
	}

	// check if the piece is reached the bottom
	if isReachedBottom := g.Piece.Container.GetBoundingBox().MaxY >= g.Board.Container.GetBoundingBox().MaxY; isReachedBottom {
		gs.lockCurrentPiece()
	}
}

func (gs *GamePlaySystem) lockCurrentPiece() {
	g := gs.game
	g.LockedPieces.Container.Merge(&g.Piece.Container)
	g.NextPiece()
}

func (gs *GamePlaySystem) clearLines() {
	removedLines := 0
	lp := &gs.game.LockedPieces
	bbox := gs.game.Board.GetBoundingBox()
	for y := bbox.MinY; y < bbox.MaxY; y++ {
		if lp.CheckLine(bbox.MinX, bbox.MaxX, y) {
			lp.RemoveLine(bbox.MinX, bbox.MaxX, y)
			lp.MoveDownBlocksByY(y)
			removedLines++
		}
	}
	if removedLines > 0 {
		earnedScore := math.Pow(2, float64(removedLines)) * 100
		gs.game.Stat.UpdateScore(gs.game.Stat.Score + int(earnedScore))
	}
}

func NewGamePlaySystem() GamePlaySystem {
	return GamePlaySystem{}
}
