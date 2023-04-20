package pr

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Position(lipgloss.Left),
		m.list.View(),
		m.viewport.View())
}
