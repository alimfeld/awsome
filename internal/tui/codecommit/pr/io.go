package pr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	tea "github.com/charmbracelet/bubbletea"
)

func (c Context) getDifferencesCmd() tea.Cmd {
	return func() tea.Msg {
		targets := c.PullRequest.PullRequestTargets[0]
		output, err := c.Client.GetDifferences(
			context.TODO(),
			&codecommit.GetDifferencesInput{
				RepositoryName:        c.Repository.RepositoryName,
				AfterCommitSpecifier:  targets.SourceReference,
				BeforeCommitSpecifier: targets.DestinationReference,
			},
		)
		if err != nil {
			return err
		}
		return differencesMsg{
			differences: output.Differences,
		}
	}
}

type differencesMsg struct {
	differences []types.Difference
}

func (c Context) getCommentsCmd() tea.Cmd {
	return func() tea.Msg {
		output, err := c.Client.GetCommentsForPullRequest(
			context.TODO(),
			&codecommit.GetCommentsForPullRequestInput{
				PullRequestId: c.PullRequest.PullRequestId,
			},
		)
		if err != nil {
			return err
		}
		return commentsMsg{
			comments: output.CommentsForPullRequestData,
		}
	}
}

type commentsMsg struct {
	comments []types.CommentsForPullRequest
}
