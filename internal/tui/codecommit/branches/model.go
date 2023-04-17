package branches

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	"github.com/samber/lo"
)

func New(client *codecommit.Client, context Context) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Branches"
	list.DisableQuitKeybindings()
	return model{
		client:  client,
		context: context,
		list:    list,
	}
}

type Context struct {
	Repository types.RepositoryMetadata
}

type model struct {
	client  *codecommit.Client
	context Context
	list    list.Model
}

func (m model) items(branches []string) []list.Item {
	return lo.Map(branches, func(branch string, _ int) list.Item {
		return item{
			name:      branch,
			isDefault: *m.context.Repository.DefaultBranch == branch,
		}
	})
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
