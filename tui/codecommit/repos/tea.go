package repos

import (
	"awsome/bubbles"
	"awsome/tui/codecommit/branches"
	"awsome/tui/codecommit/prs"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(m.list.StartSpinner(), repoNipsCmd(m.client))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case repoNipsMsg:
		m.repoNips = msg.repoNips
		return m, reposCmd(m.client, msg.repoNips)
	case reposChunkMsg:
		lo.ForEach(msg.repos, func(r types.RepositoryMetadata, _ int) {
			m.repos[*r.RepositoryName] = r
		})
		if len(m.repoNips) == len(m.repos) {
			cmd := m.list.SetItems(lo.Map(
				m.repoNips,
				func(repoNip types.RepositoryNameIdPair, _ int) list.Item {
					return item{repo: m.repos[*repoNip.RepositoryName]}
				}))
			m.list.StopSpinner()
			return m, cmd
		}
		return m, nil
	case bubbles.BodySizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "b":
				repo := m.list.SelectedItem().(item).repo
				return m, bubbles.PushModelCmd(
					branches.New(m.client, repo, m.width, m.height),
					*repo.RepositoryName)
			case "p":
				repoName := m.list.SelectedItem().(item).repo.RepositoryName
				return m, bubbles.PushModelCmd(
					prs.New(m.client, repoName, m.width, m.height),
					*repoName)
			}
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}
