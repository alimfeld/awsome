package pipeline

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	var status types.PipelineExecutionStatus
	statusStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).Width(m.viewport.Width)

	if m.pipelineExecutionSummary == nil {
		status = ""
	} else {
		status = m.pipelineExecutionSummary.Status
		switch status {
		case types.PipelineExecutionStatusSucceeded:
			statusStyle = statusStyle.
				Background(lipgloss.Color("10")).
				Foreground(lipgloss.Color("0"))
		case types.PipelineExecutionStatusFailed:
			statusStyle = statusStyle.
				Background(lipgloss.Color("9")).
				Foreground(lipgloss.Color("0"))
		case types.PipelineExecutionStatusInProgress:
			statusStyle = statusStyle.
				Background(lipgloss.Color("11")).
				Foreground(lipgloss.Color("0"))
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Position(lipgloss.Left),
		statusStyle.Render(fmt.Sprintf(">> %s <<", status)),
		m.viewport.View(),
	)
}
