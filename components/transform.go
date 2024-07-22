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

func (transform *Transform) TranformPosition(x, y int) (ax int, ay int) {
	// the origin of rotation is always (0,0)
	switch transform.rotate {
	case ROTATE_0:
		ax, ay = x, y
	case ROTATE_90:
		ax, ay = y, x
	case ROTATE_180:
		ax, ay = x, -y
	case ROTATE_270:
		ax, ay = -y, x
	}
	ax += transform.translateX
	ay += transform.translateY
	return
}
