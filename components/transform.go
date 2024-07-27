package components

const (
	ROTATE_0 = iota
	ROTATE_90
	ROTATE_180
	ROTATE_270
)

type Transform struct {
	translateX, translateY int
	rotate                 int
	reversed               bool
}

func (transform *Transform) Translate(x, y int) {
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

func (transform *Transform) TranformPosition(x, y int) (ax, ay int) {
	if transform.reversed {
		ax, ay = translatePosition(x, y, -transform.translateX, -transform.translateY)
		ax, ay = rotatePosition(ax, ay, -transform.rotate)
	} else {
		ax, ay = rotatePosition(x, y, transform.rotate)
		ax, ay = translatePosition(ax, ay, transform.translateX, transform.translateY)
	}
	return
}

func translatePosition(x, y, tx, ty int) (ax, ay int) {
	ax, ay = x+tx, y+ty
	return
}

func rotatePosition(x, y int, rotate int) (ax, ay int) {
	// the origin of rotation is always (0,0)
	rotate = ((rotate % 4) + 4) % 4
	switch rotate {
	case ROTATE_0:
		ax, ay = x, y
	case ROTATE_90:
		ax, ay = -y, x
	case ROTATE_180:
		ax, ay = -x, -y
	case ROTATE_270:
		ax, ay = y, -x
	}
	return
}

func (transform *Transform) ReversePosition(x, y int) (ax, ay int) {
	rt := transform.GetReverse()
	ax, ay = rt.TranformPosition(x, y)
	return
}
