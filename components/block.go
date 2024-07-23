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
	OriginX float64
	OriginY float64
	Width   float64
	Height  float64
	Color   Color
}
