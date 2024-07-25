package components

import "math"

type Container struct {
	CenterX   float64
	CenterY   float64
	Children  []Block
	Transform Transform
}

func (container *Container) GetChildAbsoluteBoundingBox(block *Block) BoundingBox {
	bb := block.GetBoundingBox()
	minX, minY := container.Transform.TranformPosition(bb.MinX, bb.MinY)
	maxX, maxY := container.Transform.TranformPosition(bb.MaxX, bb.MaxY)
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	minX, minY = minX+container.CenterX, minY+container.CenterY
	maxX, maxY = maxX+container.CenterX, maxY+container.CenterY
	return BoundingBox{
		MinX:   minX,
		MinY:   minY,
		MaxX:   maxX,
		MaxY:   maxY,
		Width:  maxX - minX,
		Height: maxY - minY,
	}
}

func (container *Container) GetBoundingBox() BoundingBox {
	var minX, minY, maxX, maxY float64
	for i, block := range container.Children {
		bb := block.GetBoundingBox()
		if i == 0 {
			minX, minY, maxX, maxY = bb.MinX, bb.MinY, bb.MaxX, bb.MaxY
		} else {
			minX = math.Min(minX, bb.MinX)
			minY = math.Min(minY, bb.MinY)
			maxX = math.Max(maxX, bb.MaxX)
			maxY = math.Max(maxY, bb.MaxY)
		}
	}
	// the bbox is still a rectangle even after rotation
	minX, minY = container.Transform.TranformPosition(minX, minY)
	maxX, maxY = container.Transform.TranformPosition(maxX, maxY)
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	minX, minY, maxX, maxY = minX+container.CenterX, minY+container.CenterY, maxX+container.CenterX, maxY+container.CenterY
	return BoundingBox{
		MinX:   minX,
		MinY:   minY,
		MaxX:   maxX,
		MaxY:   maxY,
		Width:  maxX - minX,
		Height: maxY - minY,
	}
}

func (container *Container) Collides(other *Container) bool {
	bb1 := container.GetBoundingBox()
	bb2 := other.GetBoundingBox()
	return bb1.Collides(&bb2)
}

func (container *Container) Contain(other *Container) bool {
	bb1 := container.GetBoundingBox()
	bb2 := other.GetBoundingBox()
	return bb1.Contain(&bb2)
}
