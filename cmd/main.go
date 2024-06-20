package main

import (
	"fmt"
	"math/rand"

	"cli-wordle/internal/game"
	"cli-wordle/words"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	wordsList := words.GetValidWords()
	solution := wordsList[rand.Intn(len(wordsList))]

	p := tea.NewProgram(game.NewModel(wordsList, solution))

	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
