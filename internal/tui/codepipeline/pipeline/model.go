package pipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
)

func New(client *codepipeline.Client, context Context) model {
	m := model{
		client:  client,
		context: context,
	}
	return m
}

type model struct {
	client   *codepipeline.Client
	context  Context
	pipeline types.PipelineDeclaration
}

type Context struct {
	Pipeline types.PipelineSummary
}
