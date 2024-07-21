package entities

import (
	"go-tetris/components"
)

type Piece struct {
	components.Container
}

func NewPiece() Piece {
	return Piece{Container: components.Container{}}
}
