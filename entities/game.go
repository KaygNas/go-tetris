package entities

type Game struct {
	Instructions Instruction
	Board        Board
	Piece        Piece
	LockedPieces LockedPieces
}

func (g *Game) NextPiece() {
	g.Piece = newPiece()
	g.movePieceToBoardCenter()
}

func (g *Game) movePieceToBoardCenter() {
	b, p := &g.Board, &g.Piece
	bbox := b.Container.GetBoundingBox()
	p.MoveToCenter((bbox.MinX+bbox.MaxX)/2, bbox.MinY)
}

func NewGame() Game {
	g := Game{
		Instructions: newInstruction(),
		Board:        newBoard(),
		Piece:        newPiece(),
		LockedPieces: newLockedPieces(),
	}
	g.movePieceToBoardCenter()
	return g
}
