package random

import "testing"

func TestToken(t *testing.T) {
	token := Token()

	if len(token) != DEFAULT_TOKEN_LENGTH {
		t.Errorf("Client ID length should be %d, got %d", DEFAULT_TOKEN_LENGTH, len(token))
	}
}
