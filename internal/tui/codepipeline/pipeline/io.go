package pipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func getPipelineDeclarationCmd(client *codepipeline.Client, pipelineName *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.GetPipeline(
			context.TODO(),
			&codepipeline.GetPipelineInput{
				Name: pipelineName,
			})
		if err != nil {
			return err
		}
		return pipelineDeclarationMsg{
			pipelineDeclaration: output.Pipeline,
		}
	}
}

type pipelineDeclarationMsg struct {
	pipelineDeclaration *types.PipelineDeclaration
}

func getPipelineExecutionCmd(client *codepipeline.Client, pipelineName *string) tea.Cmd {
	return func() tea.Msg {

		output, err := client.ListPipelineExecutions(
			context.TODO(),
			&codepipeline.ListPipelineExecutionsInput{
				PipelineName: pipelineName,
				MaxResults:   lo.ToPtr(int32(1)),
			})

		if err != nil {
			return err
		}

		var pipelineExecutionSummary *types.PipelineExecutionSummary
		if len(output.PipelineExecutionSummaries) > 0 {
			pipelineExecutionSummary = &output.PipelineExecutionSummaries[0]
		}

		return pipelineExecutionMsg{
			pipelineName,
			pipelineExecutionSummary,
		}
	}
}

type pipelineExecutionMsg struct {
	pipelineName             *string
	pipelineExecutionSummary *types.PipelineExecutionSummary
}

func getActionExecutionsCmd(client *codepipeline.Client, pipelineName *string, pipelineExecutionId *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.ListActionExecutions(
			context.TODO(),
			&codepipeline.ListActionExecutionsInput{
				PipelineName: pipelineName,
				Filter: &types.ActionExecutionFilter{
					PipelineExecutionId: pipelineExecutionId,
				},
			})

		if err != nil {
			return err
		}

		actionsExecutionDetails := lo.KeyBy(
			output.ActionExecutionDetails,
			func(a types.ActionExecutionDetail) string {
				return *a.ActionName
			},
		)

		return actionExecutionsMsg{
			pipelineName,
			pipelineExecutionId,
			actionsExecutionDetails,
		}
	}
}

type actionExecutionsMsg struct {
	pipelineName            *string
	pipelineExecutionId     *string
	actionsExecutionDetails map[string]types.ActionExecutionDetail
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
		return pipelineExecutionStartMsg{
			pipelineExecutionId: output.PipelineExecutionId,
		}
	}
}

type pipelineExecutionStartMsg struct {
	pipelineExecutionId *string
}
