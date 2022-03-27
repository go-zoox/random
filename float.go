package random

import (
	"math/rand"
)

func Float(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
