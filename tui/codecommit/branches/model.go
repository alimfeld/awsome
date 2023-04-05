package branches

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
)

func New(client *codecommit.Client, repo types.RepositoryMetadata, width, height int) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	list.Title = "Branches"
	list.DisableQuitKeybindings()
	return model{
		client: client,
		list:   list,
		repo:   repo,
		width:  width,
		height: height,
	}
}

type model struct {
	client        *codecommit.Client
	list          list.Model
	repo          types.RepositoryMetadata
	width, height int
}

func newItem(repo types.RepositoryMetadata, branch string) item {
	return item{
		name:      branch,
		isDefault: *repo.DefaultBranch == branch,
	}
}

type item struct {
	name      string
	isDefault bool
}

func (i item) Title() string { return i.name }
func (i item) Description() string {
	if i.isDefault {
		return "Default"
	}
	return ""
}
func (i item) FilterValue() string { return i.Title() }
