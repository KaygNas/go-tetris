package entities

import (
	"go-tetris/components"
	"slices"
)

type LockedPieces struct {
	components.Container
}

func buildBlockMatrix(blocks []components.Block, offsetX, offsetY, width, height int) [][]int {
	martix := make([][]int, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		for j := 0; j < width; j++ {
			row[j] = -1
		}
		martix[i] = row
	}

	for i, block := range blocks {
		martix[block.Y+offsetY][block.X+offsetX] = i
	}

	return martix
}

// check if the children fill the bbox
func (lp *LockedPieces) CheckLine(xStart, xEnd, y int) bool {
	if xStart >= xEnd {
		xStart, xEnd = xEnd, xStart
	}

	bbox := lp.Container.GetBoundingBox()
	offsetX, offsetY := -bbox.MinX, -bbox.MinY
	martix := buildBlockMatrix(lp.Blocks, offsetX, offsetY, bbox.Width, bbox.Height)

	rowIndex := y + offsetY
	if rowIndex >= len(martix) || rowIndex < 0 {
		return false
	}

	row := martix[rowIndex]
	for x := xStart; x < xEnd; x++ {
		if colIndex := x + offsetX; colIndex >= len(row) || colIndex < 0 || row[colIndex] < 0 {
			return false
		}
	}
	return true
}

// remove the line from the locked pieces
func (lp *LockedPieces) RemoveLine(xStart, xEnd, y int) {
	if xStart >= xEnd {
		xStart, xEnd = xEnd, xStart
	}

	bbox := lp.Container.GetBoundingBox()
	offsetX, offsetY := -bbox.MinX, -bbox.MinY
	martix := buildBlockMatrix(lp.Blocks, offsetX, offsetY, bbox.Width, bbox.Height)

	rowIndex := y + offsetY
	if rowIndex >= len(martix) || rowIndex < 0 {
		return
	}

	row := martix[rowIndex]
	startColIndex := xStart + offsetX
	endColIndex := xEnd + offsetX
	if startColIndex >= len(row) || startColIndex < 0 || endColIndex > len(row) || endColIndex < 0 {
		return
	}

	removalIndices := row[startColIndex:endColIndex]
	newBlocks := make([]components.Block, 0, xEnd-xStart)
	for blockIndex, block := range lp.Container.Blocks {
		if !slices.Contains(removalIndices, blockIndex) {
			newBlocks = append(newBlocks, block)
		}
	}
	lp.Container.Blocks = newBlocks
}

func (lp *LockedPieces) MoveDownBlocksByY(y int) {
	for i := range lp.Blocks {
		block := &lp.Blocks[i]
		if block.Y < y {
			block.Y += 1
		}
	}
}

func newLockedPieces() LockedPieces {
	return LockedPieces{}
}
