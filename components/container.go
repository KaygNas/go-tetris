package components

type Container struct {
	OriginX   int
	OriginY   int
	Children  []Block
	Transform Transform
}

func (container *Container) GetChildAbsPosition(block *Block) (int, int) {
	x, y := container.OriginX, container.OriginY
	dx, dy := container.Transform.TranformPosition(block.OriginX, block.OriginY)
	return x + dx, y + dy
}
