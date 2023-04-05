package prs

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	"github.com/samber/lo"
)

func New(client *codecommit.Client, context Context) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "PRs"
	list.DisableQuitKeybindings()
	m := model{
		client:  client,
		context: context,
		prs:     make(map[string]types.PullRequest),
		list:    list,
	}
	return m
}

type Context struct {
	Repository types.RepositoryMetadata
}

type model struct {
	client  *codecommit.Client
	context Context
	prIds   []string
	prs     map[string]types.PullRequest
	list    list.Model
}

func (m model) pr() types.PullRequest {
	return m.list.SelectedItem().(item).pr
}

func (m model) items() []list.Item {
	return lo.Map(m.prIds, func(prId string, _ int) list.Item {
		return item{pr: m.prs[prId]}
	})
}

type item struct {
	pr types.PullRequest
}

func (i item) Title() string {
	return fmt.Sprintf("%s: %s", *i.pr.PullRequestId, *i.pr.Title)
}
func (i item) Description() string {
	if i.pr.Description == nil {
		return ""
	}
	return *i.pr.Description
}
func (i item) FilterValue() string { return i.Title() }
