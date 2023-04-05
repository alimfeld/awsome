package prs

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
)

func New(client *codecommit.Client, repoName *string, width, height int) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	list.Title = "PRs"
	list.DisableQuitKeybindings()
	m := model{
		client:   client,
		list:     list,
		repoName: repoName,
		prs:      make(map[string]types.PullRequest),
		width:    width,
		height:   height,
	}
	return m
}

type model struct {
	client        *codecommit.Client
	list          list.Model
	repoName      *string
	prIds         []string
	prs           map[string]types.PullRequest
	width, height int
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
