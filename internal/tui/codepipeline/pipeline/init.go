package pipeline

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.context.getPipelineDeclarationCmd(),
		m.context.getPipelineExecutionCmd(),
	)
}
