package pr

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/viewport"
)

func New(client *codecommit.Client, context Context) model {
	m := model{
		client:   client,
		context:  context,
		viewport: viewport.New(10, 10),
	}
	return m
}

type Context struct {
	Repository  types.RepositoryMetadata
	PullRequest types.PullRequest
}

type model struct {
	client      *codecommit.Client
	context     Context
	differences []types.Difference
	comments    []types.CommentsForPullRequest
	viewport    viewport.Model
}
