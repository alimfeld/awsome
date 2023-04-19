package pr

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

func New(context Context) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Differences"
	m := model{
		context:  context,
		list:     list,
		viewport: viewport.New(0, 0),
	}
	return m
}

type Context struct {
	Client      *codecommit.Client
	Repository  types.RepositoryMetadata
	PullRequest types.PullRequest
}

type model struct {
	context     Context
	differences []types.Difference
	comments    []types.CommentsForPullRequest
	list        list.Model
	viewport    viewport.Model
}

type item struct {
	difference types.Difference
}

func (i item) Title() string {
	if i.difference.AfterBlob != nil {
		return *i.difference.AfterBlob.Path
	}
	return *i.difference.BeforeBlob.Path
}

func (i item) Description() string {
	return string(i.difference.ChangeType)
}

func (i item) FilterValue() string { return i.Title() }
