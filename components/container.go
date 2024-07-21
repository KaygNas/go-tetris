package components

type Container struct {
	OriginX   int
	OriginY   int
	Children  []Block
	Transform Transform
}

func (container *Container) GetChildAbsPosition(block *Block) (int, int) {
	x, y := container.OriginX, container.OriginY
	x += container.Transform.translateX
	y += container.Transform.translateY
	x += block.OriginX
	y += block.OriginY
	return x, y
}
