package pipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
)

func New(context Context) model {
	m := model{
		context: context,
		watch:   false,
		keys: keyMap{
			run: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "run"),
			),
			watch: key.NewBinding(
				key.WithKeys("w"),
				key.WithHelp("w", "toggle watch"),
			),
			back: key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back"),
			),
		},
		help:     help.New(),
		viewport: viewport.New(0, 0),
	}
	return m
}

type model struct {
	context                  Context
	watch                    bool
	keys                     keyMap
	help                     help.Model
	viewport                 viewport.Model
	width, height            int
	pipelineDeclaration      *types.PipelineDeclaration
	pipelineExecutionSummary *types.PipelineExecutionSummary
	actionExecutionDetails   map[string]types.ActionExecutionDetail
}

type Context struct {
	Client          *codepipeline.Client
	PipelineSummary types.PipelineSummary
}

type keyMap struct {
	run   key.Binding
	watch key.Binding
	back  key.Binding
}
