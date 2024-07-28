package entities

import (
	"go-tetris/components"
)

type Board struct {
	components.Container
}

func Map[T, U any](arr []T, f func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func newBoard() Board {
	width := 12
	height := 24
	blocks := make([]components.Block, 0, width*height)

	color := components.HexColor(0xCCCCCC)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			blocks = append(blocks, components.Block{
				X:     x,
				Y:     y,
				Color: color,
			})
		}
	}

	return Board{Container: components.Container{
		Blocks: blocks,
	}}
}
