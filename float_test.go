package random

import "testing"

func TestFloat(t *testing.T) {
	repeat(1000, func(i int) {
		min := float64(i)
		max := min + 100
		v := Float(max, min)
		if v < min || v > max {
			t.Error("Int() should return a number between min and max")
		}
	})
}
