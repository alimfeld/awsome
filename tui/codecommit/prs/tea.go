package prs

import (
	"awsome/bubbles"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.list.StartSpinner(),
		prIdsCmd(m.client, m.repoName),
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
			cmd := m.list.SetItems(lo.Map(m.prIds, func(prId string, _ int) list.Item {
				return item{pr: m.prs[prId]}
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
			case "esc":
				return m, bubbles.PopModelCmd()
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
