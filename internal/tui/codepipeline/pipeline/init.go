package pipeline

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	name := m.context.PipelineSummary.Name
	return tea.Batch(
		getPipelineDeclarationCmd(m.client, name),
		getPipelineExecutionCmd(m.client, name),
	)
}
