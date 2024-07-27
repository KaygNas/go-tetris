package components

import "math"

type Container struct {
	CenterX   float64
	CenterY   float64
	Children  []Block
	Transform Transform
}

func (container *Container) GetAboslutePosition(x, y float64) (ax, ay float64) {
	ax, ay = container.Transform.TranformPosition(x, y)
	ax += container.CenterX
	ay += container.CenterY
	return
}

func (container *Container) GetLocalPosition(x, y float64) (lx, ly float64) {
	lx, ly = container.Transform.ReversePosition(x, y)
	lx = lx - container.CenterX
	ly = ly - container.CenterY
	return
}

func (container *Container) GetChildAbsoluteBoundingBox(block *Block) BoundingBox {
	bb := block.GetBoundingBox()
	minX, minY := container.GetAboslutePosition(bb.MinX, bb.MinY)
	maxX, maxY := container.GetAboslutePosition(bb.MaxX, bb.MaxY)
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
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
	minX, minY = container.GetAboslutePosition(minX, minY)
	maxX, maxY = container.GetAboslutePosition(maxX, maxY)
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
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

func (container *Container) ChildrenCollide(other *Container) bool {
	for _, block := range container.Children {
		bb1 := container.GetChildAbsoluteBoundingBox(&block)
		for _, otherBlock := range other.Children {
			bb2 := other.GetChildAbsoluteBoundingBox(&otherBlock)
			if bb1.Collides(&bb2) {
				return true
			}
		}
	}
	return false
}

func (container *Container) Merge(other *Container) {
	for _, block := range other.Children {
		x, y := other.GetAboslutePosition(block.CenterX, block.CenterY)
		x, y = container.GetLocalPosition(x, y)
		b := Block{
			CenterX: x,
			CenterY: y,
			Width:   block.Width,
			Height:  block.Height,
			Color:   block.Color,
		}
		container.Children = append(container.Children, b)
	}
}
