package entities

import (
	"go-tetris/components"
)

type Board struct {
	components.Container
}

func NewBoard() Board {
	Width := 12.0
	Height := 24.0
	return Board{Container: components.Container{
		CenterX: Width / 2,
		CenterY: Height / 2,
		Children: []components.Block{
			{Width: Width, Height: Height, Color: components.HexColor(0xFFFFFF)},
		},
	}}
}
