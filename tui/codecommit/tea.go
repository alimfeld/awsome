package codecommit

import (
	"awsome/bubbles"
	"awsome/tui/codecommit/repos"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return bubbles.PushModelCmd(repos.New(m.client, m.width, m.height), "Repositories")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return ""
}
