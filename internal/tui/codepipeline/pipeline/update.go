package pipeline

import (
	"awsome/internal/core"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case pipelineDeclarationMsg:
		m.pipelineDeclaration = msg.pipelineDeclaration
		m.viewport.SetContent(m.render())
		return m, nil

	case pipelineExecutionMsg:
		if msg.pipelineName != m.context.PipelineSummary.Name {
			// the msg doesn't target this pipeline
			return m, nil
		}

		var cmds []tea.Cmd

		if m.pipelineExecutionSummary == nil ||
			m.pipelineExecutionSummary.PipelineExecutionId != msg.pipelineExecutionSummary.PipelineExecutionId ||
			m.pipelineExecutionSummary.LastUpdateTime != msg.pipelineExecutionSummary.LastUpdateTime {
			m.pipelineExecutionSummary = msg.pipelineExecutionSummary
			m.viewport.SetContent(m.render())
			cmds = append(cmds, getActionExecutionsCmd(m.client, msg.pipelineName, msg.pipelineExecutionSummary.PipelineExecutionId))
		}

		if m.watch {
			cmds = append(cmds, tea.Tick(
				5*time.Second,
				func(t time.Time) tea.Msg {
					return getPipelineExecutionCmd(
						m.client,
						m.context.PipelineSummary.Name,
					)()
				},
			))

		}

		return m, tea.Batch(cmds...)

	case actionExecutionsMsg:
		m.actionExecutionDetails = msg.actionsExecutionDetails
		m.viewport.SetContent(m.render())
		return m, nil

	case core.BodySizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 1

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			return m, core.PopModelCmd()

		case "r":
			return m, startPipelineExecutionCmd(
				m.client,
				m.context.PipelineSummary.Name,
			)
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
