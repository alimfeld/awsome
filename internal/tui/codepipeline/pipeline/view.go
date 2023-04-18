package pipeline

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/samber/lo"
)

func (m model) View() string {
	if m.pipeline == nil {
		return "loading..."
	}

	var sb strings.Builder

	if m.execution.summary != nil {
		sb.WriteString(fmt.Sprintf(">> %s <<\n", m.execution.summary.Status))
	}

	lo.ForEach(
		m.pipeline.stages,
		func(s stage, _ int) {
			sb.WriteString(fmt.Sprintf("\n# %s\n", s.name))
			lo.ForEach(
				s.groups,
				func(g group, _ int) {
					sb.WriteString(
						fmt.Sprintf("\n%s\n",
							strings.Join(
								lo.Map(
									g.actions,
									func(a action, _ int) string {
										return fmt.Sprintf(
											"%s %s",
											m.getActionStatus(a.name),
											a.name,
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

	return sb.String()
}

func (m model) getActionStatus(action string) string {
	execution, ok := m.execution.actions[action]
	if !ok {
		return "❓"
	}
	switch execution.Status {
	case types.ActionExecutionStatusInProgress:
		return "🔄"
	case types.ActionExecutionStatusSucceeded:
		return "✅"
	case types.ActionExecutionStatusFailed:
		return "❌"
	case types.ActionExecutionStatusAbandoned:
		return "❓"
	}
	return "❓"
}
