package components

import (
	"testing"
)

func TestTransform(t *testing.T) {
	x, y := 10, 20
	ax, ay := x, y
	transform := Transform{}
	transform.Translate(10, 10)

	t.Run("should return on rotate=0", func(t *testing.T) {
		ax, ay = transform.TranformPosition(x, y)
		t.Run("should return correct Position", func(t *testing.T) {
			if ax != 20 || ay != 30 {
				t.Errorf("Transform Position is not correct: %v, %v", ax, ay)
			}
		})
		t.Run("should return correct reverse Position", func(t *testing.T) {
			if bx, by := transform.ReversePosition(ax, ay); bx != x || by != y {
				t.Errorf("Transform Position is not correct: %v, %v", bx, by)
			}
		})
	})

	t.Run("should return on rotate=1", func(t *testing.T) {
		transform.RotateCW()
		ax, ay = transform.TranformPosition(x, y)
		t.Run("should return correct Position", func(t *testing.T) {
			if ax != -10 || ay != 20 {
				t.Errorf("Transform Position is not correct: %v, %v", ax, ay)
			}
		})
		t.Run("should return correct reverse Position", func(t *testing.T) {
			if bx, by := transform.ReversePosition(ax, ay); bx != x || by != y {
				t.Errorf("Transform Position is not correct: %v, %v", bx, by)
			}
		})
	})

	t.Run("should return on rotate=2", func(t *testing.T) {
		transform.RotateCW()
		ax, ay = transform.TranformPosition(x, y)
		t.Run("should return correct Position", func(t *testing.T) {
			if ax != 0 || ay != -10 {
				t.Errorf("Transform Position is not correct: %v, %v", ax, ay)
			}
		})
		t.Run("should return correct reverse Position", func(t *testing.T) {
			if bx, by := transform.ReversePosition(ax, ay); bx != x || by != y {
				t.Errorf("Transform Position is not correct: %v, %v", bx, by)
			}
		})
	})

	t.Run("should return on rotate=3", func(t *testing.T) {
		transform.RotateCW()
		ax, ay = transform.TranformPosition(x, y)
		t.Run("should return correct Position", func(t *testing.T) {
			if ax != 30 || ay != 0 {
				t.Errorf("Transform Position is not correct: %v, %v", ax, ay)
			}
		})
		t.Run("should return correct reverse Position", func(t *testing.T) {
			if bx, by := transform.ReversePosition(ax, ay); bx != x || by != y {
				t.Errorf("Transform Position is not correct: %v, %v", bx, by)
			}
		})
	})
}
