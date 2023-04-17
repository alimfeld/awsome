package pipeline

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func (m model) View() string {
	if m.pipeline == nil {
		return "loading..."
	}

	var sb strings.Builder

	lo.ForEach(
		m.pipeline.stages,
		func(s stage, _ int) {
			sb.WriteString(fmt.Sprintf("\n# %s\n", s.name))
			lo.ForEach(
				s.groups,
				func(g group, _ int) {
					sb.WriteString(
						fmt.Sprintf("\n%s\n",
							strings.Join(
								lo.Map(
									g.actions,
									func(a action, _ int) string {
										return a.name
									},
								),
								"\n",
							),
						),
					)
				},
			)
		},
	)

	return sb.String()
}
