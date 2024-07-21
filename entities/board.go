package entities

import (
	"go-tetris/components"
)

type Board struct {
	components.Container
}

func NewBoard() Board {
	return Board{Container: components.Container{}}
}
