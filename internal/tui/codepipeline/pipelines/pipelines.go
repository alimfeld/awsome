package pipelines

import (
	"awsome/internal/core"
	"awsome/internal/tui/codepipeline/pipeline"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func New(client *codepipeline.Client) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Pipelines"
	list.DisableQuitKeybindings()
	m := model{
		client: client,
		list:   list,
	}
	return m
}

type model struct {
	client    *codepipeline.Client
	pipelines []types.PipelineSummary
	list      list.Model
}

func (m model) pipeline() types.PipelineSummary {
	return m.list.SelectedItem().(item).pipeline
}

func (m model) createItems() []list.Item {
	return lo.Map(m.pipelines, func(pipeline types.PipelineSummary, _ int) list.Item {
		return item{pipeline: pipeline}
	})
}

type item struct {
	pipeline types.PipelineSummary
}

func (i item) Title() string       { return *i.pipeline.Name }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.Title() }

func listPipelinesCmd(client *codepipeline.Client) tea.Cmd {
	return func() tea.Msg {
		var maxResults int32
		maxResults = 1000
		output, err := client.ListPipelines(context.TODO(), &codepipeline.ListPipelinesInput{
			MaxResults: &maxResults,
		})
		if err != nil {
			return err
		}
		return pipelineSummariesMsg{
			pipelines: output.Pipelines,
		}
	}
}

type pipelineSummariesMsg struct {
	pipelines []types.PipelineSummary
}

func (m model) Init() tea.Cmd {
	return listPipelinesCmd(m.client)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pipelineSummariesMsg:
		m.pipelines = msg.pipelines
		cmd := m.list.SetItems(m.createItems())
		return m, cmd
	case core.BodySizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "esc":
				return m, core.PopModelCmd()
			case "enter":
				p := m.pipeline()
				return m, core.PushModelCmd(pipeline.New(m.client, pipeline.Context{Pipeline: p}), *p.Name)
			}
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}
