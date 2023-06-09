package pr

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.context.getDifferencesCmd(),
		m.context.getCommentsCmd(),
	)
}
