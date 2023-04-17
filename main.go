package main

import (
	"awsome/internal/tui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model, err := tui.New()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
