package codecommit

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

func New(cfg aws.Config) model {
	client := codecommit.NewFromConfig(cfg)
	return model{
		client: client,
	}
}

type model struct {
	client *codecommit.Client
}
