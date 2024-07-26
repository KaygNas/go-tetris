package components

import (
	"testing"
)

func TestTransform(t *testing.T) {
	x, y := 10.0, 20.0
	transform := Transform{}
	transform.RotateCW()
	transform.Translate(10, 10)
	ax, ay := transform.TranformPosition(x, y)
	t.Run("should return correct Position", func(t *testing.T) {
		if ax != 30 || ay != 20 {
			t.Errorf("Transform Position is not correct: %v, %v", ax, ay)
		}
	})
	t.Run("should return correct reverse Position", func(t *testing.T) {
		if bx, by := transform.ReversePosition(ax, ay); bx != x || by != y {
			t.Errorf("Transform Position is not correct: %v, %v", bx, by)
		}
	})
}
