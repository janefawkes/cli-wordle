package game

import (
	"strings"

	"cli-wordle/internal/render"

	tea "github.com/charmbracelet/bubbletea"
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

func NewModel(words []string, solution string) model {
	return model{
		words:    words,
		solution: solution,
		tries:    make([]string, 0, 6),
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
	return b.String()
}
