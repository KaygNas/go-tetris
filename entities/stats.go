package entities

import (
	"fmt"
	"go-tetris/components"
)

type Stat struct {
	components.Container
	Score    int
	PlayTime int
}

func (s *Stat) UpdateScore(score int, playTime int) {
	*s = newStat(score, playTime)
}

func newStat(score int, playTime int) Stat {
	stats := []string{
		"----Stats----",
		" ",
		fmt.Sprintf(" Score: %d", score),
		" ",
		fmt.Sprintf(" Play Time: %ds", playTime),
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
		Score:    score,
		PlayTime: playTime,
		Container: components.Container{
			Blocks: blocks,
		}}
}
