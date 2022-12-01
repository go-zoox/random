package random

import (
	"math/rand"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func String(length int, letters ...string) string {
	lettersX := LETTERS
	if len(letters) > 0 && letters[0] != "" {
		lettersX = letters[0]
	}

	lettersXLength := len(lettersX)
	b := make([]byte, length)
	for i := range b {
		b[i] = lettersX[rand.Intn(lettersXLength)]
	}
	return string(b)
}
