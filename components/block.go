package components

type Color struct {
	R, G, B uint8
}

func HexColor(hex uint32) Color {
	return Color{
		R: uint8((hex >> 16) & 0xFF),
		G: uint8((hex >> 8) & 0xFF),
		B: uint8(hex & 0xFF),
	}
}

type Block struct {
	/* X, Y is the left-top conner of the Block */
	X, Y  int
	Color Color
	Char  *rune
}

func (block *Block) GetBoundingBox() BoundingBox {
	// width of the block is 1, height of the block is 1
	const Width = 1
	const Height = 1
	return BoundingBox{
		MinX:   block.X,
		MinY:   block.Y,
		MaxX:   block.X + Width,
		MaxY:   block.Y + Height,
		Width:  Width,
		Height: Height,
	}
}
