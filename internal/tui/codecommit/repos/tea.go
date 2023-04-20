package repos

import (
	"awsome/internal/core"
	"awsome/internal/tui/codecommit/branches"
	"awsome/internal/tui/codecommit/prs"

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
			cmd := m.list.SetItems(m.items())
			m.list.StopSpinner()
			return m, cmd
		}
		return m, nil
	case core.BodySizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.KeyMsg:
		if m.list.FilterState() == list.Unfiltered {
			switch msg.String() {
			case "esc":
				return m, core.PopModelCmd()
			}
		}
		if m.list.FilterState() != list.Filtering {
			switch msg.String() {
			case "b":
				repo := m.repo()
				return m, core.PushModelCmd(
					branches.New(
						m.client,
						branches.Context{Repository: m.repo()},
					),
					*repo.RepositoryName)
			case "enter":
				repoName := m.repo().RepositoryName
				return m, core.PushModelCmd(
					prs.New(
						m.client,
						prs.Context{Repository: m.repo()},
					),
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
