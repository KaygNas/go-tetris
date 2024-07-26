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
	reversed   bool
}

func (transform *Transform) Translate(x, y float64) {
	transform.translateX += x
	transform.translateY += y
}

func (transform *Transform) RotateCW() {
	transform.rotate = (transform.rotate + 1) % 4
}

func (transform *Transform) GetReverse() Transform {
	rt := Transform{
		translateX: transform.translateX,
		translateY: transform.translateY,
		rotate:     transform.rotate,
		reversed:   !transform.reversed,
	}
	return rt
}

func (transform *Transform) TranformPosition(x, y float64) (ax, ay float64) {
	if transform.reversed {
		ax, ay = translatePosition(x, y, -transform.translateX, -transform.translateY)
		ax, ay = rotatePosition(ax, ay, transform.rotate)
	} else {
		ax, ay = rotatePosition(x, y, transform.rotate)
		ax, ay = translatePosition(ax, ay, transform.translateX, transform.translateY)
	}
	return
}

func translatePosition(x, y, tx, ty float64) (ax, ay float64) {
	ax, ay = x+tx, y+ty
	return
}

func rotatePosition(x, y float64, rotate int) (ax, ay float64) {
	// the origin of rotation is always (0,0)
	switch rotate {
	case ROTATE_0:
		ax, ay = x, y
	case ROTATE_90:
		ax, ay = y, x
	case ROTATE_180:
		ax, ay = x, -y
	case ROTATE_270:
		ax, ay = -y, x
	}
	return
}

func (transform *Transform) ReversePosition(x, y float64) (ax, ay float64) {
	rt := transform.GetReverse()
	ax, ay = rt.TranformPosition(x, y)
	return
}
