package random

import (
	"testing"
)

func TestSecret(t *testing.T) {
	repeat(20, func(i int) {
		secret := Secret()
		t.Logf("secret: %s", secret)

		if len(secret) != DEFAULT_SECRET_LENGTH {
			t.Error("String() should return a string of length 10")
		}
	})
}
