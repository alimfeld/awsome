package branches

import (
	"awsome/bubbles"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Init() tea.Cmd {
	return loadCmd(m.client, m.repo.RepositoryName)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loadedMsg:
		cmd := m.list.SetItems(lo.Map(msg.branches, func(b string, _ int) list.Item {
			return newItem(m.repo, b)
		}))
		return m, cmd
	case bubbles.BodySizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "esc":
				return m, bubbles.PopModelCmd()
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
