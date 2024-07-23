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
