package random

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	repeat(10, func(i int) {
		str := String(i)
		fmt.Println("xxx:", str)
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
