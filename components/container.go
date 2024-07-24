package components

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
	minX, minY, maxX, maxY := container.OriginX, container.OriginY, container.OriginX, container.OriginY
	for _, block := range container.Children {
		x, y := container.GetChildAbsPosition(&block)
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x+block.Width > maxX {
			maxX = x + block.Width
		}
		if y+block.Height > maxY {
			maxY = y + block.Height
		}
	}
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
