package components

import "math"

type Container struct {
	OriginX   float64
	OriginY   float64
	Children  []Block
	Transform Transform
}

func (container *Container) GetChildAbsPosition(block *Block) (float64, float64) {
	x, y := container.OriginX, container.OriginY
	dx, dy := container.Transform.TranformPosition(block.OriginX, block.OriginY)
	return x + dx, y + dy
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
	minX, minY = container.Transform.TranformPosition(minX, minY)
	maxX, maxY = container.Transform.TranformPosition(maxX, maxY)
	minX, minY, maxX, maxY = minX+container.OriginX, minY+container.OriginY, maxX+container.OriginX, maxY+container.OriginY
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
