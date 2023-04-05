package tui

import (
	"awsome/bubbles"
	"awsome/tui/codecommit"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) Init() tea.Cmd {
	return bubbles.PushModelCmd(
		codecommit.New(m.cfg, m.bodyWidth, m.bodyHeight),
		"CodeCommit")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.status = ""

	switch msg := msg.(type) {

	case error:
		m.status = msg.Error()
		return m, nil

	case bubbles.PushModelMsg:
		m.models = m.models.push(msg.Model, msg.Breadcrumb)
		return m, msg.Model.Init()

	case bubbles.PopModelMsg:
		m.models = m.models.pop()
		return m, nil

	case tea.WindowSizeMsg:
		m.styles = m.styles.resizeStyles(msg.Width)
		m.bodyWidth, m.bodyHeight = m.styles.bodySize(msg.Width, msg.Height)
		return m, bubbles.BodySizeCmd(m.bodyWidth, m.bodyHeight)

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
