package pr

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Position(lipgloss.Center),
		lipgloss.NewStyle().Width(m.list.Width()).Render(m.list.View()),
		m.viewport.View())
}
