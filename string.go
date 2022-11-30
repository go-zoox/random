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

	b := make([]byte, length)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(lettersX))]
	}
	return string(b)
}
