package pipeline

import (
	"awsome/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pipelineMsg:
		m.pipeline = msg.pipeline
		return m, nil
	case core.BodySizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, core.PopModelCmd()
		}
	}
	return m, nil
}
