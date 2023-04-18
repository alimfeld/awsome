package pipeline

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	name := m.context.Pipeline.Name
	return tea.Batch(
		getPipelineCmd(m.client, name),
		getPipelineExecutionCmd(m.client, name),
	)
}
