package prs

import (
	"awsome/internal/core"
	"awsome/internal/tui/codecommit/pr"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.list.StartSpinner(),
		prIdsCmd(m.client, m.context.Repository.RepositoryName),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case prIdsMsg:
		m.prIds = msg.prIds
		m.prs = make(map[string]types.PullRequest)
		return m, prsCmd(m.client, msg.prIds)
	case prMsg:
		m.prs[*msg.pr.PullRequestId] = msg.pr
		if len(m.prIds) == len(m.prs) {
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
			case "enter":
				return m, core.PushModelCmd(
					pr.New(pr.Context{
						Client:      m.client,
						Repository:  m.context.Repository,
						PullRequest: m.pr(),
					}),
					*m.pr().PullRequestId,
				)
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
