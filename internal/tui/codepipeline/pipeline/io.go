package pipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func getPipelineCmd(client codepipeline.Client, name string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.GetPipeline(
			context.TODO(),
			&codepipeline.GetPipelineInput{
				Name: &name,
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

func getPipelineExecutionCmd(client codepipeline.Client, name string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.ListPipelineExecutions(
			context.TODO(),
			&codepipeline.ListPipelineExecutionsInput{
				PipelineName: &name,
				MaxResults:   lo.ToPtr(int32(1)),
			})
		if err != nil {
			return err
		}
		var payload *types.PipelineExecutionSummary
		if len(output.PipelineExecutionSummaries) > 0 {
			payload = &output.PipelineExecutionSummaries[0]
		}
		return pipelineExecutionMsg{
			payload: payload,
		}
	}
}

type pipelineExecutionMsg struct {
	payload *types.PipelineExecutionSummary
}
