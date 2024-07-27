package entities

type Game struct {
	Instructions Instruction
	Board        Board
	Piece        Piece
	LockedPieces LockedPieces
}

func NewGame() Game {
	return Game{
		Instructions: NewInstruction(),
		Board:        NewBoard(),
		Piece:        NewPiece(),
		LockedPieces: NewLockedPieces(),
	}
}
