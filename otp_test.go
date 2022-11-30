package random

import "testing"

func TestOTP(t *testing.T) {
	token := OTP()

	if len(token) != DEFAULT_OTP_LENGTH {
		t.Errorf("Client ID length should be %d, got %d", DEFAULT_OTP_LENGTH, len(token))
	}
}
