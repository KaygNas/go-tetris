package entities

import (
	"fmt"
	"go-tetris/components"
)

type Stat struct {
	components.Container
	Score int
}

func (s *Stat) UpdateScore(score int) {
	*s = newStat(score)
}

func newStat(score int) Stat {
	stats := []string{
		"----Stats----",
		fmt.Sprintf(" Score: %d", score),
	}
	charLen := 0
	for _, instruction := range stats {
		charLen += len(instruction)
	}

	blocks := make([]components.Block, 0, charLen)

	marginLeft := 16
	marginTop := 2
	for i, instruction := range stats {
		for j, char := range instruction {
			blocks = append(blocks, components.Block{
				X:     j + marginLeft,
				Y:     i + marginTop,
				Color: components.HexColor(0xFFFFFF),
				Char:  &char,
			})
		}
	}

	return Stat{
		Score: score,
		Container: components.Container{
			Blocks: blocks,
		}}
}
