package services

import (
	"awsome/internal/core"
	"awsome/internal/tui/codecommit/repos"
	"awsome/internal/tui/codepipeline/pipelines"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func New(cfg aws.Config) model {
	list := list.New([]list.Item{
		item{
			service:     "CodeCommit",
			description: "Managed Source Control Service",
		},
		item{
			service:     "CodePipeline",
			description: "Managed Continuous Delivery Service",
		},
	}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Services"
	list.DisableQuitKeybindings()
	m := model{
		cfg:  cfg,
		list: list,
	}
	return m
}

type model struct {
	cfg  aws.Config
	list list.Model
}

type item struct {
	service     string
	description string
}

func (m model) service() string {
	return m.list.SelectedItem().(item).service
}

func (i item) Title() string       { return i.service }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.Title() }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case core.BodySizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "enter":
				service := m.service()
				switch service {
				case "CodeCommit":
					return m, core.PushModelCmd(
						repos.New(codecommit.NewFromConfig(m.cfg)),
						"CodeCommit",
					)
				case "CodePipeline":
					return m, core.PushModelCmd(
						pipelines.New(codepipeline.NewFromConfig(m.cfg)),
						"CodePipeline",
					)
				}
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
