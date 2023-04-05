package core

import tea "github.com/charmbracelet/bubbletea"

func PushModelCmd(model tea.Model, breadcrumb string) tea.Cmd {
	return func() tea.Msg {
		return PushModelMsg{
			Model:      model,
			Breadcrumb: breadcrumb,
		}
	}
}

type PushModelMsg struct {
	Model      tea.Model
	Breadcrumb string
}

func PopModelCmd() tea.Cmd {
	return func() tea.Msg {
		return PopModelMsg{}
	}
}

type PopModelMsg struct {
}

func BodySizeCmd(size Size) tea.Cmd {
	return func() tea.Msg {
		return BodySizeMsg{size}
	}
}

type BodySizeMsg struct {
	Size Size
}
