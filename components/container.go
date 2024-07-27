package components

type Container struct {
	/* X, Y is the left-top conner of the Container */
	X, Y      int
	Blocks    []Block
	Transform Transform
}

func (container *Container) GetAboslutePosition(x, y int) (ax, ay int) {
	ax, ay = container.Transform.TranformPosition(x, y)
	ax += container.X
	ay += container.Y
	return
}

func (container *Container) GetLocalPosition(x, y int) (lx, ly int) {
	lx, ly = container.Transform.ReversePosition(x, y)
	lx = lx - container.X
	ly = ly - container.Y
	return
}

func (container *Container) GetChildAbsoluteBoundingBox(block *Block) BoundingBox {
	bb := block.GetBoundingBox()
	minX, minY := container.GetAboslutePosition(bb.MinX, bb.MinY)
	maxX, maxY := container.GetAboslutePosition(bb.MaxX, bb.MaxY)
	minX, minY, maxX, maxY = minInt(minX, maxX), minInt(minY, maxY), maxInt(maxX, minX), maxInt(maxY, minY)
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
	var minX, minY, maxX, maxY int
	for i, block := range container.Blocks {
		bb := container.GetChildAbsoluteBoundingBox(&block)
		if i == 0 {
			minX, minY, maxX, maxY = bb.MinX, bb.MinY, bb.MaxX, bb.MaxY
		} else {
			minX = minInt(minX, bb.MinX)
			minY = minInt(minY, bb.MinY)
			maxX = maxInt(maxX, bb.MaxX)
			maxY = maxInt(maxY, bb.MaxY)
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

func (container *Container) BouningBoxCollide(other *Container) bool {
	bb1 := container.GetBoundingBox()
	bb2 := other.GetBoundingBox()
	return bb1.Collides(&bb2)
}

func (container *Container) BoundingBoxContain(other *Container) bool {
	bb1 := container.GetBoundingBox()
	bb2 := other.GetBoundingBox()
	return bb1.Contain(&bb2)
}

func (container *Container) IsChildrenCollide(other *Container) bool {
	for _, block := range container.Blocks {
		bb1 := container.GetChildAbsoluteBoundingBox(&block)
		for _, otherBlock := range other.Blocks {
			bb2 := other.GetChildAbsoluteBoundingBox(&otherBlock)
			if bb1.Collides(&bb2) {
				return true
			}
		}
	}
	return false
}

func (container *Container) Merge(other *Container) {
	for _, block := range other.Blocks {
		bbox := block.GetBoundingBox()
		minX, minY := other.GetAboslutePosition(bbox.MinX, bbox.MinY)
		maxX, maxY := other.GetAboslutePosition(bbox.MaxX, bbox.MaxY)
		minX, minY = container.GetLocalPosition(minX, minY)
		maxX, maxY = container.GetLocalPosition(maxX, maxY)
		x := minInt(minX, maxX)
		y := minInt(minY, maxY)
		b := Block{
			X:     x,
			Y:     y,
			Color: block.Color,
		}
		container.Blocks = append(container.Blocks, b)
	}
}

// remove child in the container
func (container *Container) RemoveChild(block *Block) {
	for i, b := range container.Blocks {
		if b == *block {
			container.Blocks = append(container.Blocks[:i], container.Blocks[i+1:]...)
			return
		}
	}
}
