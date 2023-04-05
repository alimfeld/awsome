package pr

import (
	"awsome/core"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

func New(client *codecommit.Client, context Context, size core.Size) model {
	m := model{
		client:  client,
		context: context,
		size:    size,
	}
	return m
}

type Context struct {
	Repository  types.RepositoryMetadata
	PullRequest types.PullRequest
}

type model struct {
	client   *codecommit.Client
	context  Context
	size     core.Size
	comments []types.CommentsForPullRequest
}
