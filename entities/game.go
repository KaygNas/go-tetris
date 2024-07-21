package entities

type Game struct {
	Board Board
	Piece Piece
}

func NewGame() Game {
	return Game{
		Board: NewBoard(),
		Piece: NewPiece(),
	}
}
