package entities

import (
	"go-tetris/components"
)

type LockedPieces struct {
	components.Container
}

func NewLockedPieces() LockedPieces {
	return LockedPieces{}
}
