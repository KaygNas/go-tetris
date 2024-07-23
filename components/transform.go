package components

const (
	ROTATE_0 = iota
	ROTATE_90
	ROTATE_180
	ROTATE_270
)

type Transform struct {
	translateX float64
	translateY float64
	rotate     int
}

func (transform *Transform) Translate(x, y float64) {
	transform.translateX += x
	transform.translateY += y
}

func (transform *Transform) RotateCW() {
	transform.rotate = (transform.rotate + 1) % 4
}

func (transform *Transform) TranformPosition(x, y float64) (ax float64, ay float64) {
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
