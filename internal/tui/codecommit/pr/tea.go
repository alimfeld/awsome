package pr

import (
	"awsome/internal/core"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.getDiffCmd(),
		m.getCommentsCmd(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case getCommentsMsg:
		m.comments = msg.comments
		return m, nil
	case getDiffMsg:
		m.differences = msg.differences
		m.viewport = m.updateViewport()
		return m, nil
	case core.BodySizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, core.PopModelCmd()
		}
	}
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) updateViewport() viewport.Model {
	content := strings.Join(lo.Map(
		m.differences,
		func(difference types.Difference, _ int) string {
			return string(difference.ChangeType)
		}), "\n")
	m.viewport.SetContent(content)
	return m.viewport
}

func (m model) View() string {
	return m.viewport.View()
}
