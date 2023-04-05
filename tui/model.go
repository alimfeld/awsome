package tui

import (
	"awsome/core"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func New() (model, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return model{}, err
	}
	return model{
		cfg:    cfg,
		styles: Styles(),
	}, nil
}

type model struct {
	cfg      aws.Config
	bodySize core.Size
	styles   styles
	models   stack
	status   string
}

type stack struct {
	items []item
}

type item struct {
	model      tea.Model
	breadcrumb string
}

func (s stack) push(model tea.Model, breadcrumb string) stack {
	item := item{
		model:      model,
		breadcrumb: breadcrumb,
	}
	s.items = append(s.items, item)
	return s
}

func (s stack) pop() stack {
	s.items = s.items[:len(s.items)-1]
	return s
}

func (s stack) peek() tea.Model {
	n := len(s.items)
	if n == 0 {
		return nil
	}
	return s.items[n-1].model
}

func (s stack) update(model tea.Model) stack {
	n := len(s.items)
	s.items[n-1].model = model
	return s
}

func (s stack) breadcrumbs() []string {
	return lo.Map(s.items, func(i item, _ int) string {
		return i.breadcrumb
	})
}
