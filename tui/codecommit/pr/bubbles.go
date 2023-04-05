package pr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) getCommentsCmd() tea.Cmd {
	return func() tea.Msg {
		output, err := m.client.GetCommentsForPullRequest(
			context.TODO(),
			&codecommit.GetCommentsForPullRequestInput{
				PullRequestId: m.context.PullRequest.PullRequestId,
			},
		)
		if err != nil {
			return err
		}
		return getCommentsMsg{
			comments: output.CommentsForPullRequestData,
		}
	}
}

type getCommentsMsg struct {
	comments []types.CommentsForPullRequest
}
