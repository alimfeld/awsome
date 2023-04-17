package repos

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	"github.com/samber/lo"
)

func New(client *codecommit.Client) model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Repositories"
	list.DisableQuitKeybindings()
	m := model{
		client: client,
		repos:  make(map[string]types.RepositoryMetadata),
		list:   list,
	}
	return m
}

type model struct {
	client   *codecommit.Client
	repoNips []types.RepositoryNameIdPair
	repos    map[string]types.RepositoryMetadata
	list     list.Model
}

func (m model) repo() types.RepositoryMetadata {
	return m.list.SelectedItem().(item).repo
}

func (m model) items() []list.Item {
	return lo.Map(
		m.repoNips,
		func(repoNip types.RepositoryNameIdPair, _ int) list.Item {
			return item{repo: m.repos[*repoNip.RepositoryName]}
		})
}

type item struct {
	repo types.RepositoryMetadata
}

func (i item) Title() string { return *i.repo.RepositoryName }
func (i item) Description() string {
	if i.repo.RepositoryDescription == nil {
		return "n/a"
	}
	return *i.repo.RepositoryDescription
}
func (i item) FilterValue() string { return i.Title() }
