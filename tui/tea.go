package tui

import (
	"awsome/core"
	"awsome/tui/services"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Init() tea.Cmd {
	return core.PushModelCmd(services.New(m.cfg), "awsome")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.status = ""

	switch msg := msg.(type) {

	case error:
		m.status = msg.Error()
		return m, nil

	case core.PushModelMsg:
		m.models = m.models.push(msg.Model, msg.Breadcrumb)
		return m, tea.Batch(
			msg.Model.Init(),
			core.BodySizeCmd(m.bodyWidth, m.bodyHeight),
		)

	case core.PopModelMsg:
		m.models = m.models.pop()
		return m, core.BodySizeCmd(m.bodyWidth, m.bodyHeight)

	case tea.WindowSizeMsg:
		m.styles = m.styles.resizeStyles(msg.Width)
		m.bodyWidth, m.bodyHeight = m.styles.bodySize(msg.Width, msg.Height)
		return m, core.BodySizeCmd(m.bodyWidth, m.bodyHeight)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	delegate := m.models.peek()
	if delegate == nil {
		return m, nil
	}
	var cmd tea.Cmd
	delegate, cmd = delegate.Update(msg)
	m.models = m.models.update(delegate)
	return m, cmd
}

func (m model) View() string {
	content := ""
	delegate := m.models.peek()
	if delegate != nil {
		content += delegate.View()
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.header.Render(strings.Join(m.models.breadcrumbs(), " > ")),
		m.styles.body.Render(content),
		m.styles.footer.Render(m.status),
	)
}
