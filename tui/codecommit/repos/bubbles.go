package repos

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

func repoNipsCmd(client *codecommit.Client) tea.Cmd {
	return func() tea.Msg {
		output, err := client.ListRepositories(context.TODO(), nil)
		if err != nil {
			return err
		}
		return repoNipsMsg{repoNips: output.Repositories}
	}
}

type repoNipsMsg struct {
	repoNips []types.RepositoryNameIdPair
}

func reposCmd(client *codecommit.Client, repoNips []types.RepositoryNameIdPair) tea.Cmd {
	const chunkSize = 25
	var cmds = lo.Map(
		lo.Chunk(
			lo.Map(repoNips,
				func(r types.RepositoryNameIdPair, _ int) string {
					return *r.RepositoryName
				}),
			chunkSize,
		),
		func(names []string, _ int) tea.Cmd {
			return func() tea.Msg {
				output, err := client.BatchGetRepositories(context.TODO(),
					&codecommit.BatchGetRepositoriesInput{RepositoryNames: names})
				if err != nil {
					return err
				}
				return reposChunkMsg{repos: output.Repositories}
			}
		},
	)
	return tea.Batch(cmds...)
}

type reposChunkMsg struct {
	repos []types.RepositoryMetadata
}
