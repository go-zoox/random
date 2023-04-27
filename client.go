package random

const DEFAULT_CLIENT_ID_LENGTH = 20
const DEFAULT_CLIENT_SECRET_LENGTH = 40
const DEFAULT_CLIENT_LETTERS = "abcde0123456789"

func ClientID(length ...int) string {
	lengthX := DEFAULT_CLIENT_ID_LENGTH
	if len(length) > 0 && length[0] != 0 {
		lengthX = length[0]
	}

	return String(lengthX, DEFAULT_CLIENT_LETTERS)
}

func ClientSecret(length ...int) string {
	lengthX := DEFAULT_CLIENT_SECRET_LENGTH
	if len(length) > 0 && length[0] != 0 {
		lengthX = length[0]
	}

	return String(lengthX, DEFAULT_CLIENT_LETTERS)
}
