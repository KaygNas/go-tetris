package components

type Container struct {
	originX   int
	originY   int
	children  []Block
	transform Transform
}
