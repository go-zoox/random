package random

const DEFAULT_OTP_LENGTH = 6
const DEFAULT_OTP_LETTERS = "0123456789"

func OTP(length ...int) string {
	lengthX := DEFAULT_OTP_LENGTH
	if len(length) > 0 && length[0] != 0 {
		lengthX = length[0]
	}

	return String(lengthX, DEFAULT_OTP_LETTERS)
}
