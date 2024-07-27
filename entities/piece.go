package entities

import (
	"go-tetris/components"
	"os"
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
	PicecType       uint
	cachedTransform components.Transform
}

func saveTransform(p *Piece) {
	p.cachedTransform = p.Transform
}

func restoreTransform(p *Piece) {
	p.Transform = p.cachedTransform
}

func (p *Piece) RestoreTransform() {
	restoreTransform(p)
}

func (p *Piece) MoveLeft() {
	saveTransform(p)
	p.Container.Transform.Translate(-1, 0)
}
func (p *Piece) MoveRight() {
	saveTransform(p)
	p.Container.Transform.Translate(1, 0)
}
func (p *Piece) MoveUp() {
	saveTransform(p)
	p.Container.Transform.Translate(0, -1)
}
func (p *Piece) MoveDown() {
	saveTransform(p)
	p.Container.Transform.Translate(0, 1)
}
func (p *Piece) RotateCW() {
	saveTransform(p)
	p.Container.Transform.RotateCW()
}

func (p *Piece) MoveInto(target components.BoundingBoxer) {
	sBbox := p.GetBoundingBox()
	tBbox := target.GetBoundingBox()
	if tBbox.Contain(&sBbox) {
		return
	}

	var dx, dy int
	if sBbox.MinX < tBbox.MinX {
		dx = tBbox.MinX - sBbox.MinX
	} else if sBbox.MaxX > tBbox.MaxX {
		dx = tBbox.MaxX - sBbox.MaxX
	}
	if sBbox.MinY < tBbox.MinY {
		dy = tBbox.MinY - sBbox.MinY
	} else if sBbox.MaxY > tBbox.MaxY {
		dy = tBbox.MaxY - sBbox.MaxY
	}
	p.Transform.Translate(dx, dy)
}

func NewPieceI() Piece {
	color := components.HexColor(0x00FFFF)
	color2 := color
	if mode := os.Getenv("MODE"); mode == "development" {
		color2 = components.HexColor(0x00AAAA) // Color for development mode
	}
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: 2,
			Blocks: []components.Block{
				{X: 0, Y: -2, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color2},
				{X: 0, Y: 1, Color: color2},
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
