package random

const DEFAULT_TOKEN_LENGTH = 20
const DEFAULT_TOKEN_LETTERS = "abcdefghijklmnopqrstuvwxyz0123456789"

func Token(length ...int) string {
	lengthX := DEFAULT_TOKEN_LENGTH
	if len(length) > 0 && length[0] != 0 {
		lengthX = length[0]
	}

	return String(lengthX, DEFAULT_TOKEN_LETTERS)
}
