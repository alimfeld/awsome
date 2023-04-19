package pipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (c Context) getPipelineDeclarationCmd() tea.Cmd {
	return func() tea.Msg {
		output, err := c.Client.GetPipeline(
			context.TODO(),
			&codepipeline.GetPipelineInput{
				Name: c.PipelineSummary.Name,
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

func (c Context) getPipelineExecutionCmd() tea.Cmd {
	return func() tea.Msg {

		output, err := c.Client.ListPipelineExecutions(
			context.TODO(),
			&codepipeline.ListPipelineExecutionsInput{
				PipelineName: c.PipelineSummary.Name,
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
			c.PipelineSummary.Name,
			pipelineExecutionSummary,
		}
	}
}

type pipelineExecutionMsg struct {
	pipelineName             *string
	pipelineExecutionSummary *types.PipelineExecutionSummary
}

func (c Context) getActionExecutionsCmd(pipelineExecutionId *string) tea.Cmd {
	return func() tea.Msg {
		output, err := c.Client.ListActionExecutions(
			context.TODO(),
			&codepipeline.ListActionExecutionsInput{
				PipelineName: c.PipelineSummary.Name,
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
			c.PipelineSummary.Name,
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

func (c Context) startPipelineExecutionCmd() tea.Cmd {
	return func() tea.Msg {
		output, err := c.Client.StartPipelineExecution(
			context.TODO(),
			&codepipeline.StartPipelineExecutionInput{
				Name: c.PipelineSummary.Name,
			},
		)
		if err != nil {
			return err
		}
		return pipelineExecutionStartMsg{
			output.PipelineExecutionId,
		}
	}
}

type pipelineExecutionStartMsg struct {
	pipelineExecutionId *string
}
