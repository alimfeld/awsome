package pipeline

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/samber/lo"
)

func New(client *codepipeline.Client, context Context) model {
	m := model{
		client:   client,
		context:  context,
		watch:    true,
		viewport: viewport.New(0, 0),
	}
	return m
}

type model struct {
	client                   *codepipeline.Client
	context                  Context
	watch                    bool
	viewport                 viewport.Model
	pipelineDeclaration      *types.PipelineDeclaration
	pipelineExecutionSummary *types.PipelineExecutionSummary
	actionExecutionDetails   map[string]types.ActionExecutionDetail
}

type Context struct {
	PipelineSummary types.PipelineSummary
}

func (m model) render() string {
	if m.pipelineDeclaration == nil {
		return ""
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

	return sb.String()
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
