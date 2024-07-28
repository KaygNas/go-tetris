package entities

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
	p.MoveToCenter((bbox.MinX+bbox.MaxX)/2, bbox.MinY)
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
