package entities

import "math"

type Game struct {
	Stat         Stat
	Instructions Instruction
	Board        Board
	Piece        Piece
	LockedPieces LockedPieces
}

func (g *Game) NewGame() {
	*g = NewGame()
}

func (g *Game) NextPiece() {
	g.Piece = newPiece()
	g.movePieceToBoardCenter()
}

func (g *Game) movePieceToBoardCenter() {
	b, p := &g.Board, &g.Piece
	bbox := b.Container.GetBoundingBox()
	p.MoveToCenter((bbox.MinX+bbox.MaxX)/2, bbox.MinY-1)
}

func (g *Game) EnsureNoCollided() {
	// check if the piece is collided with the locked pieces
	if isChildrenCollided := g.LockedPieces.Container.IsChildrenCollide(&g.Piece.Container); isChildrenCollided {
		g.Piece.RestoreTransform()
	}

	// check if the piece is out of bounds
	boardBbox := g.Board.Container.GetBoundingBox()
	pieceBbox := g.Piece.Container.GetBoundingBox()
	isOutOfBounds := pieceBbox.MinX < boardBbox.MinX || pieceBbox.MaxX > boardBbox.MaxX || pieceBbox.MaxY > boardBbox.MaxY
	if isOutOfBounds {
		g.Piece.MoveInto(&g.Board.Container)
	}
}

func (g *Game) TryLockCurrentPiece() {
	// check if the piece can move down
	g.Piece.MoveDown()
	if isChildrenCollided := g.LockedPieces.Container.IsChildrenCollide(&g.Piece.Container); isChildrenCollided {
		g.Piece.RestoreTransform()
		g.lockCurrentPiece()
	} else {
		g.Piece.RestoreTransform()
	}

	// check if the piece is reached the bottom
	if isReachedBottom := g.Piece.Container.GetBoundingBox().MaxY >= g.Board.Container.GetBoundingBox().MaxY; isReachedBottom {
		g.lockCurrentPiece()
	}
}

func (g *Game) lockCurrentPiece() {
	g.LockedPieces.Container.Merge(&g.Piece.Container)
	g.NextPiece()
}

func (g *Game) ClearLines() {
	removedLines := 0
	lp := g.LockedPieces
	bbox := g.Board.GetBoundingBox()
	for y := bbox.MinY; y < bbox.MaxY; y++ {
		if lp.CheckLine(bbox.MinX, bbox.MaxX, y) {
			lp.RemoveLine(bbox.MinX, bbox.MaxX, y)
			lp.MoveDownBlocksByY(y)
			removedLines++
		}
	}
	if removedLines > 0 {
		earnedScore := math.Pow(2, float64(removedLines)) * 100
		g.Stat.UpdateScore(g.Stat.Score + int(earnedScore))
	}
}

func NewGame() Game {
	g := Game{
		Stat:         newStat(0),
		Instructions: newInstruction(),
		Board:        newBoard(),
		Piece:        newPiece(),
		LockedPieces: newLockedPieces(),
	}
	g.movePieceToBoardCenter()
	return g
}
