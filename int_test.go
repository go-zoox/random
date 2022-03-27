package random

import "testing"

func TestInt(t *testing.T) {
	repeat(1000, func(i int) {
		min := i
		max := i + 100
		v := Int(max, min)
		if v < min || v > max {
			t.Error("Int() should return a number between min and max")
		}
	})
}
