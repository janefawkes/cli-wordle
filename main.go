package main

import (
	"fmt"
	"math/rand"
	"strings"

	"cli-wordle/words"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	words    []string
	solution string
	guess    string
	tries    []string
	message  string
	won      bool
	lost     bool
}

var (
	cellStyle  = lipgloss.NewStyle().Width(3).Height(1).Align(lipgloss.Center).Padding(0, 1).Margin(0, 0)
	absent     = cellStyle.Background(lipgloss.Color("#787C7E")).Foreground(lipgloss.Color("#FFFFFF"))
	present    = cellStyle.Background(lipgloss.Color("#C9B458")).Foreground(lipgloss.Color("#FFFFFF"))
	correct    = cellStyle.Background(lipgloss.Color("#6AAA64")).Foreground(lipgloss.Color("#FFFFFF"))
	defaultBox = cellStyle.Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#000000"))
)

func main() {
	words := words.GetValidWords()

	solution := words[rand.Intn(len(words))]

	p := tea.NewProgram(model{
		words:    words,
		solution: solution,
		tries:    make([]string, 0, 6),
	})

	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(m.guess) != 5 {
				m.message = "Guess must be 5 letters"
				return m, nil
			}
			if !validGuess(m.guess, m.words) {
				m.message = "Invalid word"
				return m, nil
			}
			m.tries = append(m.tries, m.guess)
			if m.guess == m.solution {
				m.won = true
				m.message = "Congratulations! You guessed the word!"
				return m, tea.Quit
			}
			if len(m.tries) == 6 {
				m.lost = true
				m.message = fmt.Sprintf("You lost! The word was %s", m.solution)
				return m, tea.Quit
			}
			m.guess = ""
			m.message = fmt.Sprintf("Tries remaining: %d", 6-len(m.tries))
		case "backspace":
			if len(m.guess) > 0 {
				m.guess = m.guess[:len(m.guess)-1]
			}
		case "ctrl+c":
			return m, tea.Quit
		default:
			if len(msg.String()) == 1 && len(m.guess) < 5 {
				m.guess += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("Wordle CLI ⋆ ˚｡⋆\n\n")
	for _, attempt := range m.tries {
		b.WriteString(renderGuess(attempt, m.solution) + "\n")
	}
	if len(m.tries) < 6 {
		b.WriteString(renderGuess(m.guess, "") + "\n")
		for i := len(m.tries) + 1; i < 6; i++ {
			b.WriteString(renderGuess("", "") + "\n")
		}
	}
	b.WriteString("\n" + m.message + "\n")
	return b.String()
}

func renderGuess(guess, solution string) string {
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

func validGuess(guess string, words []string) bool {
	for _, word := range words {
		if guess == word {
			return true
		}
	}
	return false
}
