package components

import (
	"testing"
)

func newContainer() Container {
	return Container{
		Blocks: []Block{
			{X: 0, Y: 0},
			{X: 0, Y: 1},
		},
	}
}

func TestGetPosition(t *testing.T) {
	t.Run("should return correct absolute position", func(t *testing.T) {
		c := newContainer()
		c.X = 5
		c.Y = 10
		c.Transform.Translate(5, 5)
		if ax, ay := c.GetAboslutePosition(5, 5); ax != 15 || ay != 20 {
			t.Errorf("Absolute Position is not correct: %v, %v", ax, ay)
		}
	})

	t.Run("should return correct local position", func(t *testing.T) {
		c := newContainer()
		c.X = 5
		c.Y = 10
		c.Transform.Translate(5, 5)
		if lx, ly := c.GetLocalPosition(15, 20); lx != 5 || ly != 5 {
			t.Errorf("Local Position is not correct: %v, %v", lx, ly)
		}
	})
}

func TestGetBoundingBox(t *testing.T) {
	t.Run("should return correct bounding box", func(t *testing.T) {
		c := newContainer()

		if bb := c.GetBoundingBox(); bb.MinX != 0 || bb.MinY != 0 || bb.MaxX != 1 || bb.MaxY != 2 {
			t.Errorf("BoundingBox is not correct: %v", bb)
		}
	})

	t.Run("should return correct bounding box at not zero origin", func(t *testing.T) {
		c := newContainer()

		c.X = 10
		c.Y = 10
		if bb := c.GetBoundingBox(); bb.MinX != 10 || bb.MinY != 10 || bb.MaxX != 11 || bb.MaxY != 12 {
			t.Errorf("BoundingBox at origin (%v, %v) is not correct: %v", c.X, c.Y, bb)
		}
	})

	t.Run("should return correct bounding box after translate", func(t *testing.T) {
		c := newContainer()

		c.Transform = Transform{translateX: 10, translateY: 10}
		if bb := c.GetBoundingBox(); bb.MinX != 10 || bb.MinY != 10 || bb.MaxX != 11 || bb.MaxY != 12 {
			t.Errorf("BoundingBox translate at (%v, %v) is not correct: %v", c.Transform.translateX, c.Transform.translateY, bb)
		}
	})

	t.Run("should return correct bounding box after rotate", func(t *testing.T) {
		c := newContainer()

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -2 || bb.MinY != 0 || bb.MaxX != 0 || bb.MaxY != 1 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != -1 || bb.MinY != -2 || bb.MaxX != 0 || bb.MaxY != 0 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != 0 || bb.MinY != -1 || bb.MaxX != 2 || bb.MaxY != 0 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}

		c.Transform.RotateCW()
		if bb := c.GetBoundingBox(); bb.MinX != 0 || bb.MinY != 0 || bb.MaxX != 1 || bb.MaxY != 2 {
			t.Errorf("BoundingBox rotate at (%v) is not correct: %v", c.Transform.rotate, bb)
		}
	})
}

func TestMerge(t *testing.T) {
	t.Run("should merge correctly", func(t *testing.T) {
		c1 := newContainer()

		c2 := newContainer()
		c2.Transform.RotateCW()

		len1 := len(c1.Blocks)
		len2 := len(c2.Blocks)

		c1.Merge(&c2)

		if len(c1.Blocks) != len1+len2 {
			t.Errorf("Children length is not correct: %v", c1.Blocks)
		}

		nc1, nc2 := c1.Blocks[len1], c1.Blocks[len1+1]
		if nc1.X != -1 || nc1.Y != 0 || nc2.X != -2 || nc2.Y != 0 {
			t.Errorf("Children Position is not correct: %v", c1.Blocks)
		}
	})
}
