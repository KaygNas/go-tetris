package components

import (
	"fmt"
)

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY, Width, Height int
}

type BoundingBoxer interface {
	GetBoundingBox() BoundingBox
}

func (bb BoundingBox) String() string {
	intToString := func(n int) string {
		return fmt.Sprintf("%v", n)
	}
	return "BoundingBox{MinX: " + intToString(bb.MinX) + ", MinY: " + intToString(bb.MinY) + ", MaxX: " + intToString(bb.MaxX) + ", MaxY: " + intToString(bb.MaxY) + ", Width: " + intToString(bb.Width) + ", Height: " + intToString(bb.Height) + "}"
}

func (bb *BoundingBox) Collides(other *BoundingBox) bool {
	return bb.MinX < other.MaxX && bb.MaxX > other.MinX && bb.MinY < other.MaxY && bb.MaxY > other.MinY
}

func (bb *BoundingBox) Contain(other *BoundingBox) bool {
	return bb.MinX <= other.MinX && bb.MaxX >= other.MaxX && bb.MinY <= other.MinY && bb.MaxY >= other.MaxY
}
