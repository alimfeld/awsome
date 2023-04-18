package pipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func getPipelineCmd(client *codepipeline.Client, name *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.GetPipeline(
			context.TODO(),
			&codepipeline.GetPipelineInput{
				Name: name,
			})
		if err != nil {
			return err
		}
		return pipelineMsg{
			payload: output.Pipeline,
		}
	}
}

type pipelineMsg struct {
	payload *types.PipelineDeclaration
}

func getPipelineExecutionCmd(client *codepipeline.Client, name *string) tea.Cmd {
	return func() tea.Msg {

		pipelineExeuctions, err := client.ListPipelineExecutions(
			context.TODO(),
			&codepipeline.ListPipelineExecutionsInput{
				PipelineName: name,
				MaxResults:   lo.ToPtr(int32(1)),
			})

		if err != nil {
			return err
		}

		if len(pipelineExeuctions.PipelineExecutionSummaries) == 0 {
			return pipelineExecutionMsg{}
		}

		summary := &pipelineExeuctions.PipelineExecutionSummaries[0]
		actionExecutions, err := client.ListActionExecutions(
			context.TODO(),
			&codepipeline.ListActionExecutionsInput{
				PipelineName: name,
				Filter: &types.ActionExecutionFilter{
					PipelineExecutionId: summary.PipelineExecutionId,
				},
			})

		if err != nil {
			return err
		}

		return pipelineExecutionMsg{
			summary: summary,
			actions: lo.KeyBy(
				actionExecutions.ActionExecutionDetails,
				func(a types.ActionExecutionDetail) string {
					return *a.ActionName
				},
			),
		}
	}
}

type pipelineExecutionMsg struct {
	summary *types.PipelineExecutionSummary
	actions map[string]types.ActionExecutionDetail
}

func startPipelineExecutionCmd(client *codepipeline.Client, name *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.StartPipelineExecution(
			context.TODO(),
			&codepipeline.StartPipelineExecutionInput{
				Name: name,
			},
		)
		if err != nil {
			return err
		}
		return pipelineStartMsg{
			id: output.PipelineExecutionId,
		}
	}
}

type pipelineStartMsg struct {
	id *string
}
