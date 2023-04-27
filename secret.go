package random

const DEFAULT_SECRET_LENGTH = 20
const DEFAULT_SECRET_LETTERS = "abcdefghjklmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789123456789123456789"

func Secret(length ...int) string {
	lengthX := DEFAULT_SECRET_LENGTH
	if len(length) > 0 && length[0] != 0 {
		lengthX = length[0]
	}

	return String(lengthX, DEFAULT_SECRET_LETTERS)
}
