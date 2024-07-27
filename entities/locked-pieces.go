package entities

import (
	"go-tetris/components"
	"math"
	"sort"
)

type LockedPieces struct {
	components.Container
}

// check if the children fill the bbox
func (lp *LockedPieces) CheckFullInBoundingBox(bb *components.BoundingBox) bool {
	var childrenInBbox = make([]components.Block, int(math.Floor(bb.Width)))
	// collect all children in the bbox
	for _, b := range lp.Children {
		bb2 := b.GetBoundingBox()
		if bb.Contain(&bb2) {
			childrenInBbox = append(childrenInBbox, b)
		}
	}

	// sort the children by minX
	sort.Slice(childrenInBbox, func(i, j int) bool {
		return childrenInBbox[i].GetBoundingBox().MinX < childrenInBbox[j].GetBoundingBox().MinX
	})
	// find the longest continuous width
	var continuousWidth float64
	var lastX float64
	for _, b := range childrenInBbox {
		bb2 := b.GetBoundingBox()
		if bb2.MinX == lastX {
			continuousWidth += bb2.Width
		} else {
			continuousWidth = bb2.Width
		}
		lastX = bb2.MinX
	}

	return continuousWidth == bb.Width
}

func (lp *LockedPieces) RemoveChildrenInBoundingBox(bb1 *components.BoundingBox) {
	for _, b := range lp.Children {
		bb2 := b.GetBoundingBox()
		if bb1.Contain(&bb2) {
			lp.RemoveChild(&b)
		}
	}
}

func NewLockedPieces() LockedPieces {
	return LockedPieces{}
}
