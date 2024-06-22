package game

import (
	"strings"

	"cli-wordle/internal/render"
	"cli-wordle/internal/util"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	words        []string
	solution     string
	guess        string
	tries        []string
	message      string
	won          bool
	lost         bool
	letterStatus map[rune]string
}

func NewModel(words []string, solution string) model {
	return model{
		words:        words,
		solution:     solution,
		tries:        make([]string, 0, 6),
		letterStatus: make(map[rune]string),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("Wordle CLI ⋆ ˚｡⋆\n\n")
	for _, attempt := range m.tries {
		b.WriteString(render.RenderGuess(attempt, m.solution) + "\n")
	}
	if len(m.tries) < 6 {
		b.WriteString(render.RenderGuess(m.guess, "") + "\n")
		for i := len(m.tries) + 1; i < 6; i++ {
			b.WriteString(render.RenderGuess("", "") + "\n")
		}
	}
	b.WriteString("\n" + m.message + "\n")
	b.WriteString(render.RenderKeyboard(m.letterStatus) + "\n")
	return b.String()
}

func validGuess(guess string, words []string) bool {
	for _, word := range words {
		if guess == word {
			return true
		}
	}
	return false
}

func updateLetterStatus(m *model, guess, solution string) {
	guess = strings.ToUpper(guess)
	solution = strings.ToUpper(solution)
	solutionCounts := make(map[rune]int)

	for _, letter := range solution {
		solutionCounts[letter]++
	}

	for i, letter := range guess {
		if solution[i] == uint8(letter) {
			m.letterStatus[letter] = util.Correct
			solutionCounts[letter]--
		}
	}

	for i, letter := range guess {
		if solution[i] != uint8(letter) {
			if solutionCounts[letter] > 0 {
				if m.letterStatus[letter] != util.Correct {
					m.letterStatus[letter] = util.Present
				}
				solutionCounts[letter]--
			} else {
				if m.letterStatus[letter] != util.Correct && m.letterStatus[letter] != util.Present {
					m.letterStatus[letter] = util.Absent
				}
			}
		}
	}
}
