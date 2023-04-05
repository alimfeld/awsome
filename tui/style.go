package tui

import (
	"awsome/core"

	"github.com/charmbracelet/lipgloss"
)

type styles struct {
	header lipgloss.Style
	body   lipgloss.Style
	footer lipgloss.Style
}

func Styles() styles {
	return styles{
		header: lipgloss.NewStyle(),
		body:   lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true),
		footer: lipgloss.NewStyle(),
	}
}

func (s styles) bodySize(windowSize core.Size) core.Size {
	hv := s.header.GetVerticalFrameSize()
	bh, bv := s.body.GetFrameSize()
	fv := s.footer.GetVerticalFrameSize()
	return core.Size{
		Width:  windowSize.Width - bh,
		Height: windowSize.Height - hv - bv - fv - 2,
	}
}

func (s styles) resizeStyles(width int) styles {
	return styles{
		header: s.header.Width(width - s.header.GetHorizontalFrameSize()),
		body:   s.body.Width(width - s.body.GetHorizontalFrameSize()),
		footer: s.footer.Width(width - s.footer.GetHorizontalFrameSize()),
	}
}
