package random

import (
	"math/rand"
)

func Int(max, min int) int {
	return min + rand.Intn(max-min)
}
