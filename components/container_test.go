package components

import (
	"testing"
)

func newContainer() Container {
	return Container{
		Children: []Block{
			{CenterX: 0, CenterY: 0, Width: 10, Height: 10},
			{CenterX: 0, CenterY: 10, Width: 10, Height: 10},
		},
	}
}

func TestGetBoundingBox(t *testing.T) {
	t.Run("should return correct bounding box", func(t *testing.T) {
		c := newContainer()

		if bb := c.GetBoundingBox(); bb.MinX != -5 || bb.MinY != -5 || bb.MaxX != 5 || bb.MaxY != 15 {
			t.Errorf("BoundingBox is not correct: %v", bb)
		}
	})

	t.Run("should return correct bounding box at not zero origin", func(t *testing.T) {
		c := newContainer()

		c.OriginX = 10
		c.OriginY = 10
		if bb := c.GetBoundingBox(); bb.MinX != 5 || bb.MinY != 5 || bb.MaxX != 15 || bb.MaxY != 25 {
			t.Errorf("BoundingBox at origin (%v, %v) is not correct: %v", c.OriginX, c.OriginY, bb)
		}
	})

	t.Run("should return correct bounding box after translate", func(t *testing.T) {
		c := newContainer()

		c.Transform = Transform{translateX: 10, translateY: 10}
		if bb := c.GetBoundingBox(); bb.MinX != 5 || bb.MinY != 5 || bb.MaxX != 15 || bb.MaxY != 25 {
			t.Errorf("BoundingBox translate at (%v, %v) is not correct: %v", c.Transform.translateX, c.Transform.translateY, bb)
		}
	})

	t.Run("should return correct bounding box after rotate", func(t *testing.T) {
		c := newContainer()

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -5 || bb.MinY != -5 || bb.MaxX != 15 || bb.MaxY != 5 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -5 || bb.MinY != -15 || bb.MaxX != 5 || bb.MaxY != 5 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -15 || bb.MinY != -5 || bb.MaxX != 5 || bb.MaxY != 5 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -5 || bb.MinY != -5 || bb.MaxX != 5 || bb.MaxY != 15 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}
	})
}
