package entities

type Game struct {
	board Board
	piece Piece
}

func NewGame() Game {
	return Game{
		board: NewBoard(),
		piece: NewPiece(),
	}
}
