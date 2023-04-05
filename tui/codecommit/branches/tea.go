package branches

import (
	"awsome/core"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return loadCmd(m.client, m.context.Repository.RepositoryName)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loadedMsg:
		cmd := m.list.SetItems(m.items(msg.branches))
		return m, cmd
	case core.BodySizeMsg:
		m.size = msg.Size
		m.list.SetSize(msg.Size.Width, msg.Size.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "esc":
				return m, core.PopModelCmd()
			}
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}
