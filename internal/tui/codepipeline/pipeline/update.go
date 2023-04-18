package pipeline

import (
	"awsome/internal/core"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pipelineMsg:
		m.setPipeline(*msg.payload)
		return m, nil
	case pipelineExecutionMsg:
		m.execution = execution{
			msg.summary,
			msg.actions,
		}
		return m, tea.Tick(
			5*time.Second,
			func(t time.Time) tea.Msg {
				return getPipelineExecutionCmd(
					m.client,
					m.context.Pipeline.Name,
				)()
			},
		)
	case core.BodySizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, core.PopModelCmd()
		case "r":
			m.client.StartPipelineExecution(
				context.TODO(),
				&codepipeline.StartPipelineExecutionInput{
					Name: m.context.Pipeline.Name,
				},
			)
			return m, nil
		}
	}
	return m, nil
}
