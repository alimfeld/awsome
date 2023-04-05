package codecommit

import (
	"awsome/core"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func New(cfg aws.Config, size core.Size) model {
	client := codecommit.NewFromConfig(cfg)
	return model{
		client: client,
		size:   size,
	}
}

type model struct {
	client *codecommit.Client
	size   core.Size
}
