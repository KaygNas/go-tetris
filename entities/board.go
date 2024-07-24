package entities

import (
	"go-tetris/components"
)

type Board struct {
	components.Container
}

func NewBoard() Board {
	return Board{Container: components.Container{
		OriginX: 0,
		OriginY: 0,
		Children: []components.Block{
			{Width: 12, Height: 24, Color: components.HexColor(0xFFFFFF)},
		},
	}}
}
