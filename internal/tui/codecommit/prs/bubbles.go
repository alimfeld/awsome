package prs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func prIdsCmd(client *codecommit.Client, repoName *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.ListPullRequests(context.TODO(),
			&codecommit.ListPullRequestsInput{
				RepositoryName:    repoName,
				PullRequestStatus: types.PullRequestStatusEnumOpen,
			})
		if err != nil {
			return err
		}
		return prIdsMsg{prIds: output.PullRequestIds}
	}
}

type prIdsMsg struct {
	prIds []string
}

func prsCmd(client *codecommit.Client, prIds []string) tea.Cmd {
	return tea.Batch(lo.Map(prIds, func(prId string, _ int) tea.Cmd {
		return func() tea.Msg {
			output, err := client.GetPullRequest(context.TODO(),
				&codecommit.GetPullRequestInput{
					PullRequestId: &prId,
				})
			if err != nil {
				return err
			}
			return prMsg{pr: *output.PullRequest}
		}

	})...)
}

type prMsg struct {
	pr types.PullRequest
}
