package render

import (
	"strings"

	"cli-wordle/internal/styles"
	"cli-wordle/internal/util"
)

var (
	cellStyle  = styles.CellStyle
	absent     = styles.Absent
	present    = styles.Present
	correct    = styles.Correct
	defaultBox = styles.DefaultBox
)

func RenderGuess(guess, solution string) string {
	var b strings.Builder

	if len(guess) < 5 || len(solution) < 5 {
		for i := 0; i < 5; i++ {
			if i < len(guess) {
				b.WriteString(defaultBox.Render(string(guess[i])))
			} else {
				b.WriteString(defaultBox.Render(" "))
			}
		}
		return b.String()
	}

	solutionCounts := make(map[rune]int)
	guessCounts := make(map[rune]int)

	for _, letter := range solution {
		solutionCounts[letter]++
	}

	for i := 0; i < 5; i++ {
		if guess[i] == solution[i] {
			solutionCounts[rune(guess[i])]--
			guessCounts[rune(guess[i])]++
		}
	}

	for i := 0; i < 5; i++ {
		letter := rune(guess[i])
		switch {
		case solution == "":
			b.WriteString(defaultBox.Render(string(letter)))
		case guess[i] == solution[i]:
			b.WriteString(correct.Render(string(letter)))
		case solutionCounts[letter] > 0 && guessCounts[letter] < solutionCounts[letter]:
			b.WriteString(present.Render(string(letter)))
			guessCounts[letter]++
		default:
			b.WriteString(absent.Render(string(letter)))
		}
	}

	return b.String()
}

func RenderKeyboard(letterStatus map[rune]string) string {
	keys := "QWERTYUIOP\nASDFGHJKL\nZXCVBNM"
	var b strings.Builder

	for _, r := range keys {
		if r == '\n' {
			b.WriteRune('\n')
			continue
		}
		status, exists := letterStatus[r]
		if !exists {
			b.WriteString(defaultBox.Render(string(r)))
		} else {
			switch status {
			case util.Absent:
				b.WriteString(absent.Render(string(r)))
			case util.Present:
				b.WriteString(present.Render(string(r)))
			case util.Correct:
				b.WriteString(correct.Render(string(r)))
			}
		}
	}
	return b.String()
}
