package pipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return m.getPipelineCmd()
}

func (m model) getPipelineCmd() tea.Cmd {
	return func() tea.Msg {
		output, err := m.client.GetPipeline(context.TODO(), &codepipeline.GetPipelineInput{
			Name: m.context.Pipeline.Name,
		})
		if err != nil {
			return err
		}
		return pipelineMsg{
			pipeline: *output.Pipeline,
		}
	}
}

type pipelineMsg struct {
	pipeline types.PipelineDeclaration
}
