package entities

import (
	"go-tetris/components"
)

type Instruction struct {
	components.Container
}

func newInstruction() Instruction {
	instructions := []string{
		"----Control----",
		" ",
		" ← Left",
		" → Right",
		" ↓ Down",
		" ↑ Rotate",
		" ",
		"----------------",
		" ",
		" Press Q to Quit",
	}
	charLen := 0
	for _, instruction := range instructions {
		charLen += len(instruction)
	}

	blocks := make([]components.Block, 0, charLen)

	marginLeft := 16
	marginTop := 2
	for i, instruction := range instructions {
		for j, char := range instruction {
			blocks = append(blocks, components.Block{
				X:     j + marginLeft,
				Y:     i + marginTop,
				Color: components.HexColor(0xFFFF00),
				Char:  &char,
			})
		}
	}

	return Instruction{Container: components.Container{
		Blocks: blocks,
	}}
}
