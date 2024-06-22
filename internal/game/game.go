package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

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
			updateLetterStatus(&m, m.guess, m.solution)
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
