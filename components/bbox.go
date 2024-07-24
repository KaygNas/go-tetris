package components

type BoundingBox struct {
	MinX, MinY, MaxX, MaxY, Width, Height float64
}

type BoundingBoxer interface {
	GetBoundingBox() BoundingBox
}

func (bb *BoundingBox) Collides(other *BoundingBox) bool {
	return bb.MinX < other.MaxX && bb.MaxX > other.MinX && bb.MinY < other.MaxY && bb.MaxY > other.MinY
}

func (bb *BoundingBox) Contain(other *BoundingBox) bool {
	return bb.MinX <= other.MinX && bb.MaxX >= other.MaxX && bb.MinY <= other.MinY && bb.MaxY >= other.MaxY
}
