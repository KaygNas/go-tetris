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

func TestGetPosition(t *testing.T) {
	t.Run("should return correct absolute position", func(t *testing.T) {
		c := newContainer()
		c.CenterX = 5
		c.CenterY = 10
		c.Transform.Translate(5, 5)
		if ax, ay := c.GetAboslutePosition(5, 5); ax != 15 || ay != 20 {
			t.Errorf("Absolute Position is not correct: %v, %v", ax, ay)
		}
	})

	t.Run("should return correct local position", func(t *testing.T) {
		c := newContainer()
		c.CenterX = 5
		c.CenterY = 10
		c.Transform.Translate(5, 5)
		if lx, ly := c.GetLocalPosition(15, 20); lx != 5 || ly != 5 {
			t.Errorf("Local Position is not correct: %v, %v", lx, ly)
		}
	})
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

		c.CenterX = 10
		c.CenterY = 10
		if bb := c.GetBoundingBox(); bb.MinX != 5 || bb.MinY != 5 || bb.MaxX != 15 || bb.MaxY != 25 {
			t.Errorf("BoundingBox at origin (%v, %v) is not correct: %v", c.CenterX, c.CenterY, bb)
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

func TestMerge(t *testing.T) {
	t.Run("should merge correctly", func(t *testing.T) {
		c1 := newContainer()
		c1.CenterX = 5
		c1.CenterY = 5

		c2 := newContainer()
		c2.CenterX = 10
		c2.CenterY = 10

		len1 := len(c1.Children)
		len2 := len(c2.Children)

		c1.Merge(&c2)

		if len(c1.Children) != len1+len2 {
			t.Errorf("Children length is not correct: %v", c1.Children)
		}

		nc1, nc2 := c1.Children[len1], c1.Children[len1+1]
		if nc1.CenterX != 5 || nc1.CenterY != 5 || nc2.CenterX != 5 || nc2.CenterY != 15 {
			t.Errorf("Children Position is not correct: %v", c1.Children)
		}
	})
}
