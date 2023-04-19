package pr

import (
	"awsome/internal/core"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case commentsMsg:
		m.comments = msg.comments
		return m, nil

	case differencesMsg:
		m.differences = msg.differences
		m.updateList()
		return m, nil

	case core.BodySizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, core.PopModelCmd()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *model) updateList() {
	m.list.SetItems(
		lo.Map(m.differences, func(difference types.Difference, _ int) list.Item {
			return item{difference}
		}))
}
