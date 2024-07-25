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
	CenterX float64
	CenterY float64
	Width   float64
	Height  float64
	Color   Color
}

func (block *Block) GetBoundingBox() BoundingBox {
	halfWidth := block.Width / 2
	halfHeight := block.Height / 2
	return BoundingBox{
		MinX:   block.CenterX - halfWidth,
		MinY:   block.CenterY - halfHeight,
		MaxX:   block.CenterX + halfWidth,
		MaxY:   block.CenterY + halfHeight,
		Width:  block.Width,
		Height: block.Height,
	}
}
