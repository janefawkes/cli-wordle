package util

func ValidGuess(guess string, words []string) bool {
	for _, word := range words {
		if guess == word {
			return true
		}
	}
	return false
}
