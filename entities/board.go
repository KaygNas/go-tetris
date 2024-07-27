package entities

import (
	"go-tetris/components"
)

type Board struct {
	components.Container
}

func NewBoard() Board {
	color := components.HexColor(0xFFFFFF)
	width := 12
	height := 24
	blocks := make([]components.Block, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			blocks[y*width+x] = components.Block{
				X:     x,
				Y:     y,
				Color: color,
			}
		}
	}

	return Board{Container: components.Container{
		Blocks: blocks,
	}}
}
