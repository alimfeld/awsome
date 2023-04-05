package pr

import (
	"awsome/core"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return m.getCommentsCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case getCommentsMsg:
		m.comments = msg.comments
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, core.PopModelCmd()
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("%v", m.comments)
}
