package components

import (
	"testing"
)

func TestGetBoundingBox(t *testing.T) {
	newContainer := func() Container {
		return Container{
			Children: []Block{
				{Width: 10, Height: 10},
				{OriginX: 0, OriginY: 10, Width: 10, Height: 10},
			},
		}
	}

	c := newContainer()
	if bb := c.GetBoundingBox(); bb.MinX != 0 || bb.MinY != 0 || bb.MaxX != 10 || bb.MaxY != 20 {
		t.Errorf("BoundingBox is not correct: %v", bb)
	}

	c.OriginX = 10
	c.OriginY = 10
	if bb := c.GetBoundingBox(); bb.MinX != 10 || bb.MinY != 10 || bb.MaxX != 20 || bb.MaxY != 30 {
		t.Errorf("BoundingBox at origin (%v, %v) is not correct: %v", c.OriginX, c.OriginY, bb)
	}

	c.Transform = Transform{translateX: 10, translateY: 10}
	if bb := c.GetBoundingBox(); bb.MinX != 20 || bb.MinY != 20 || bb.MaxX != 30 || bb.MaxY != 40 {
		t.Errorf("BoundingBox translate at (%v, %v) is not correct: %v", c.Transform.translateX, c.Transform.translateY, bb)
	}

	c.Transform.RotateCW()
	if bb := c.GetBoundingBox(); bb.MinX != 20 || bb.MinY != 20 || bb.MaxX != 40 || bb.MaxY != 30 {
		t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
	}
}
