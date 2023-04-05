package codecommit

import (
	"awsome/core"
	"awsome/tui/codecommit/repos"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return core.PushModelCmd(repos.New(m.client, m.size), "Repositories")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return ""
}
