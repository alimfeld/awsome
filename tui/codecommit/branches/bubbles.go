package branches

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	tea "github.com/charmbracelet/bubbletea"
)

func loadCmd(client *codecommit.Client, repoName *string) tea.Cmd {
	return func() tea.Msg {
		output, err := client.ListBranches(context.TODO(), &codecommit.ListBranchesInput{RepositoryName: repoName})
		if err != nil {
			return err
		}
		return loadedMsg{
			branches: output.Branches,
		}
	}
}

type loadedMsg struct {
	branches []string
}
