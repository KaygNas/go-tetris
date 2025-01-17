package entities

import (
	"go-tetris/components"
	"math/rand"
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

func (p *Piece) MoveToCenter(tx, ty int) {
	bbox := p.GetBoundingBox()
	cx, cy := (bbox.MinX+bbox.MaxX)/2, (bbox.MinY+bbox.MaxY)/2
	dx := tx - cx
	dy := ty - cy
	p.Transform.Translate(dx, dy)
}

func NewPieceI() Piece {
	color := components.HexColor(0x00FFFF)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -2,
			Blocks: []components.Block{
				{X: 0, Y: -2, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: 0, Y: 1, Color: color},
			},
		}}
}

func NewPieceJ() Piece {
	color := components.HexColor(0xFF0000)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: 0, Y: -2, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: -1, Y: 0, Color: color},
			},
		}}
}

func NewPieceL() Piece {
	color := components.HexColor(0x00FF00)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: 0, Y: -2, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: 1, Y: 0, Color: color},
			},
		}}
}

func NewPieceO() Piece {
	color := components.HexColor(0x0000FF)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: 0, Y: -1, Color: color},
				{X: 1, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: 1, Y: 0, Color: color},
			},
		}}
}

func NewPieceS() Piece {
	color := components.HexColor(0xFFFF00)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: 1, Y: -1, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: -1, Y: 0, Color: color},
			},
		}}
}

func NewPieceZ() Piece {
	color := components.HexColor(0xFF00FF)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: -1, Y: -1, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
				{X: 1, Y: 0, Color: color},
			},
		}}
}

func NewPieceT() Piece {
	color := components.HexColor(0xFFAA00)
	return Piece{
		PicecType: PIECE_TYPE_I,
		Container: components.Container{
			X: 0,
			Y: -1,
			Blocks: []components.Block{
				{X: -1, Y: -1, Color: color},
				{X: 0, Y: -1, Color: color},
				{X: 1, Y: -1, Color: color},
				{X: 0, Y: 0, Color: color},
			},
		}}
}

func newPiece() Piece {
	pieceType := rand.Int() % 7
	switch pieceType {
	case PIECE_TYPE_I:
		return NewPieceI()
	case PIECE_TYPE_J:
		return NewPieceJ()
	case PIECE_TYPE_L:
		return NewPieceL()
	case PIECE_TYPE_O:
		return NewPieceO()
	case PIECE_TYPE_S:
		return NewPieceS()
	case PIECE_TYPE_T:
		return NewPieceT()
	case PIECE_TYPE_Z:
		return NewPieceZ()
	default:
		return NewPieceI()
	}
}
