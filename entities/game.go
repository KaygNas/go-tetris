package entities

type Game struct {
	Board        Board
	Piece        Piece
	LockedPieces LockedPieces
}

func NewGame() Game {
	return Game{
		Board:        NewBoard(),
		Piece:        NewPiece(),
		LockedPieces: NewLockedPieces(),
	}
}
