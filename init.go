package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UpdateSeed() {
	rand.Seed(time.Now().UnixNano())
}
