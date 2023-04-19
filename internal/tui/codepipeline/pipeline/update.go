package pipeline

import (
	"awsome/internal/core"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case pipelineDeclarationMsg:
		m.pipelineDeclaration = msg.pipelineDeclaration
		m.updateContent()
		return m, nil

	case pipelineExecutionMsg:
		if msg.pipelineName != m.context.PipelineSummary.Name {
			// the msg doesn't target this pipeline
			return m, nil
		}

		var cmds []tea.Cmd

		if m.pipelineExecutionSummary == nil ||
			m.pipelineExecutionSummary.PipelineExecutionId !=
				msg.pipelineExecutionSummary.PipelineExecutionId ||
			m.pipelineExecutionSummary.LastUpdateTime !=
				msg.pipelineExecutionSummary.LastUpdateTime {
			m.pipelineExecutionSummary = msg.pipelineExecutionSummary
			m.updateContent()
			cmds = append(cmds, m.context.getActionExecutionsCmd(
				msg.pipelineExecutionSummary.PipelineExecutionId))
		}

		if m.watch {
			cmds = append(cmds, tea.Tick(
				5*time.Second,
				func(t time.Time) tea.Msg {
					return m.context.getPipelineExecutionCmd()()
				},
			))

		}

		return m, tea.Batch(cmds...)

	case actionExecutionsMsg:
		m.actionExecutionDetails = msg.actionsExecutionDetails
		m.updateContent()
		return m, nil

	case core.BodySizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.back):
			return m, core.PopModelCmd()

		case key.Matches(msg, m.keys.run):
			m.watch = true
			return m, tea.Sequence(
				m.context.startPipelineExecutionCmd(),
				m.context.getPipelineExecutionCmd(),
			)

		case key.Matches(msg, m.keys.watch):
			m.watch = !m.watch
			var cmd tea.Cmd
			if m.watch {
				cmd = m.context.getPipelineExecutionCmd()
			}
			return m, cmd
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m *model) updateContent() {
	if m.pipelineDeclaration == nil {
		return
	}

	var sb strings.Builder

	lo.ForEach(
		m.pipelineDeclaration.Stages,
		func(s types.StageDeclaration, _ int) {

			sb.WriteString(fmt.Sprintf("\n# %s\n", *s.Name))

			actionsByRunOrder := lo.GroupBy(s.Actions,
				func(action types.ActionDeclaration) int {
					return int(*action.RunOrder)
				},
			)
			var runOrders []int
			for r := range actionsByRunOrder {
				runOrders = append(runOrders, r)
			}
			sort.Ints(runOrders)

			lo.ForEach(
				runOrders,
				func(runOrder int, _ int) {
					sb.WriteString(
						fmt.Sprintf("\n%s\n",
							strings.Join(
								lo.Map(
									actionsByRunOrder[runOrder],
									func(a types.ActionDeclaration, _ int) string {
										return fmt.Sprintf(
											"%s %s",
											renderActionStatus(m.actionExecutionDetails[*a.Name].Status),
											*a.Name,
										)
									},
								),
								"\n",
							),
						),
					)
				},
			)
		},
	)

	m.viewport.SetContent(sb.String())
}

func renderActionStatus(status types.ActionExecutionStatus) string {
	switch status {
	case types.ActionExecutionStatusInProgress:
		return ">"
	case types.ActionExecutionStatusAbandoned:
		return "_"
	case types.ActionExecutionStatusSucceeded:
		return "✅"
	case types.ActionExecutionStatusFailed:
		return "❌"
	}
	return " "
}
