package codecommit

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

type model struct {
	client        *codecommit.Client
	width, height int
}

func New(cfg aws.Config, width, height int) model {
	client := codecommit.NewFromConfig(cfg)
	return model{
		client: client,
		width:  width,
		height: height,
	}
}
