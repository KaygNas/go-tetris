package components

const (
	ROTATE_0 = iota
	ROTATE_90
	ROTATE_180
	ROTATE_270
)

type Transform struct {
	translateX int
	translateY int
	rotate     int
}

func (transform *Transform) Translate(x, y int) {
	transform.translateX += x
	transform.translateY += y
}

func (transform *Transform) RotateCW() {
	transform.rotate = (transform.rotate + 1) % 4
}
