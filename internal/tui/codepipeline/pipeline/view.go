package pipeline

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	status := m.renderStatus()
	help := m.renderHelp()

	m.viewport.Width = m.width
	m.viewport.Height = m.height - lipgloss.Height(status) - lipgloss.Height(help)

	return lipgloss.JoinVertical(
		lipgloss.Position(lipgloss.Left),
		status,
		m.viewport.View(),
		help,
	)
}

func (m model) renderStatus() string {
	style := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Width(m.width)
	var status types.PipelineExecutionStatus
	if m.pipelineExecutionSummary == nil {
		status = ""
	} else {
		status = m.pipelineExecutionSummary.Status
		switch status {
		case types.PipelineExecutionStatusSucceeded:
			style = style.
				Background(lipgloss.Color("10")).
				Foreground(lipgloss.Color("0"))
		case types.PipelineExecutionStatusFailed:
			style = style.
				Background(lipgloss.Color("9")).
				Foreground(lipgloss.Color("0"))
		case types.PipelineExecutionStatusInProgress:
			style = style.
				Background(lipgloss.Color("11")).
				Foreground(lipgloss.Color("0"))
		}
	}
	if m.watch {
		return style.Render(fmt.Sprintf("... %s ...", status))
	} else {
		return style.Render(fmt.Sprintf(">>> %s <<<", status))
	}
}

func (m model) renderHelp() string {
	m.help.Width = m.width
	if m.watch {
		m.keys.watch.SetHelp("w", "unwatch")
	} else {
		m.keys.watch.SetHelp("w", "watch")
	}
	return m.help.View(m.keys)
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.run, k.watch}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.run, k.watch},
	}
}
