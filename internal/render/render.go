package render

import (
	"strings"

	"cli-wordle/internal/styles"
)

func RenderGuess(guess, solution string) string {
	var b strings.Builder

	if len(guess) < 5 || len(solution) < 5 {
		for i := 0; i < 5; i++ {
			if i < len(guess) {
				b.WriteString(styles.DefaultBox.Render(string(guess[i])))
			} else {
				b.WriteString(styles.DefaultBox.Render(" "))
			}
		}
		return b.String()
	}

	solutionCounts := make(map[rune]int)
	guessCounts := make(map[rune]int)

	for _, letter := range solution {
		solutionCounts[letter]++
	}

	// First pass: mark correctly positioned letters (green)
	for i := 0; i < 5; i++ {
		if guess[i] == solution[i] {
			solutionCounts[rune(guess[i])]--
			guessCounts[rune(guess[i])]++
		}
	}

	// Second pass: mark each letter appropriately
	for i := 0; i < 5; i++ {
		letter := rune(guess[i])
		switch {
		case solution == "":
			b.WriteString(styles.DefaultBox.Render(string(letter)))
		case guess[i] == solution[i]:
			b.WriteString(styles.Correct.Render(string(letter)))
		case solutionCounts[letter] > 0 && guessCounts[letter] < solutionCounts[letter]:
			b.WriteString(styles.Present.Render(string(letter)))
			guessCounts[letter]++
		default:
			b.WriteString(styles.Absent.Render(string(letter)))
		}
	}

	return b.String()
}
