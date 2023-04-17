package pipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

func (m model) View() string {
	return lipgloss.JoinVertical(lipgloss.Position(lipgloss.Center),
		lo.Map(m.pipeline.Stages,
			func(stage types.StageDeclaration, _ int) string {
				actions := lo.Map(stage.Actions,
					func(action types.ActionDeclaration, _ int) string {
						return lipgloss.
							NewStyle().
							Border(lipgloss.RoundedBorder(), true).
							Render(*action.Name)
					})

				return lipgloss.
					NewStyle().
					Border(lipgloss.NormalBorder(), true).
					Render(
						lipgloss.JoinVertical(lipgloss.Center,
							append([]string{*stage.Name}, actions...)...,
						),
					)
			},
		)...)
}
