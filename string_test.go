package random

import "testing"

func TestString(t *testing.T) {
	repeat(1000, func(i int) {
		str := String(i)
		if len(str) != i {
			t.Error("String() should return a string of length 10")
		}
	})
}

func repeat(times int, f func(i int)) {
	for i := 0; i < times; i++ {
		f(i)
	}
}
