package pipeline

import (
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/samber/lo"
)

func New(client *codepipeline.Client, context Context) model {
	m := model{
		client:  client,
		context: context,
	}
	return m
}

type model struct {
	client  *codepipeline.Client
	context Context
	*pipeline
	execution
}

type Context struct {
	Pipeline types.PipelineSummary
}

type pipeline struct {
	name   string
	stages []stage
}

type stage struct {
	name   string
	groups []group
}

type group struct {
	order   int
	actions []action
}

type action struct {
	name string
}

type execution struct {
	summary *types.PipelineExecutionSummary
	actions map[string]types.ActionExecutionDetail
}

func (m *model) setPipeline(p types.PipelineDeclaration) {
	m.pipeline = &pipeline{
		name: *p.Name,
		stages: lo.Map(
			p.Stages,
			func(s types.StageDeclaration, _ int) stage {
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
				return stage{
					name: *s.Name,
					groups: lo.Map(
						runOrders,
						func(runOrder int, _ int) group {
							return group{
								order: runOrder,
								actions: lo.Map(
									actionsByRunOrder[runOrder],
									func(a types.ActionDeclaration, _ int) action {
										return action{
											name: *a.Name,
										}
									},
								),
							}
						},
					),
				}
			},
		),
	}
}
