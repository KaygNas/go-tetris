package entities

import (
	"go-tetris/components"
)

// Piece Type for the Tetris Piece
const (
	PIECE_TYPE_I = iota
	PIECE_TYPE_J
	PIECE_TYPE_L
	PIECE_TYPE_O
	PIECE_TYPE_S
	PIECE_TYPE_T
	PIECE_TYPE_Z
)

type Piece struct {
	components.Container
	PicecType uint
}

func (p *Piece) MoveLeft() {
	p.Container.Transform.Translate(-1, 0)
}
func (p *Piece) MoveRight() {
	p.Container.Transform.Translate(1, 0)
}
func (p *Piece) MoveDown() {
	p.Container.Transform.Translate(0, 1)
}
func (p *Piece) RotateCW() {
	p.Container.Transform.RotateCW()
}

func NewPieceI() Piece {
	color := components.HexColor(0x00FFFF)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			OriginX: 0.5,
			OriginY: -2,
			Children: []components.Block{
				{OriginX: -0.5, OriginY: -1.5, Width: 1, Height: 1, Color: color},
				{OriginX: -0.5, OriginY: -0.5, Width: 1, Height: 1, Color: color},
				{OriginX: -0.5, OriginY: 0.5, Width: 1, Height: 1, Color: color},
				{OriginX: -0.5, OriginY: 1.5, Width: 1, Height: 1, Color: color},
			},
		}}
}

func NewPiece() Piece {
	pieceType := PIECE_TYPE_I
	switch pieceType {
	case PIECE_TYPE_I:
		return NewPieceI()
	default:
		return NewPieceI()
	}
}
