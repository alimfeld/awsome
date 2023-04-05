package repos

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
)

type model struct {
	client        *codecommit.Client
	list          list.Model
	repoNips      []types.RepositoryNameIdPair
	repos         map[string]types.RepositoryMetadata
	width, height int
}

type item struct {
	repo types.RepositoryMetadata
}

func (i item) Title() string       { return *i.repo.RepositoryName }
func (i item) Description() string { return *i.repo.RepositoryDescription }
func (i item) FilterValue() string { return i.Title() }

func New(client *codecommit.Client, width, height int) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	list.Title = "Repositories"
	list.DisableQuitKeybindings()
	m := model{
		client: client,
		list:   list,
		repos:  make(map[string]types.RepositoryMetadata),
		width:  width,
		height: height,
	}
	return m
}
