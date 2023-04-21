package diff

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
)

var (
	insertStyle lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("10")).Foreground(lipgloss.Color("0"))
	deleteStyle lipgloss.Style = lipgloss.NewStyle().Background(lipgloss.Color("9")).Foreground(lipgloss.Color("0"))
)

type Model struct {
	before, after []byte
}

func New(before, after []byte) Model {
	return Model{
		before, after,
	}
}

func (m Model) View() string {
	edits := myers.ComputeEdits("", string(m.before), string(m.after))
	unified := gotextdiff.ToUnified("before", "after", string(m.before), edits)
	diff := fmt.Sprint(unified)
	var sb strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(diff))
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "+"):
			sb.WriteString(insertStyle.Render(line))
		case strings.HasPrefix(line, "-"):
			sb.WriteString(deleteStyle.Render(line))
		default:
			sb.WriteString(line)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
