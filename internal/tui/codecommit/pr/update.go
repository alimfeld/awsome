package pr

import (
	"awsome/internal/core"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/samber/lo"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case commentsMsg:
		m.comments = msg.comments
		return m, nil

	case differencesMsg:
		m.differences = msg.differences
		m.updateList()
		cmds := m.updateSelectedDifference()
		return m, tea.Batch(cmds...)

	case blobMsg:
		m.blobsCache[msg.id] = msg.content
		if m.isBlobOfSelectedDifference(msg.id) {
			m.updateViewport()
		}
		return m, nil

	case core.BodySizeMsg:
		listHeight := msg.Height / 4
		m.list.SetSize(msg.Width, listHeight)
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - listHeight
		return m, nil

	case tea.KeyMsg:
		if m.list.FilterState() == list.Unfiltered {
			switch msg.String() {
			case "esc":
				return m, core.PopModelCmd()
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	var cmds = m.updateSelectedDifference()
	return m, tea.Batch(append(cmds, cmd)...)
}

func (m *model) updateList() {
	m.list.SetItems(
		lo.Map(m.differences,
			func(difference types.Difference, _ int) list.Item {
				return item{difference}
			}))
}

func (m *model) updateSelectedDifference() []tea.Cmd {
	var cmds []tea.Cmd

	selection := m.list.SelectedItem().(item).difference
	changed := m.selectedDifference == nil || selection != *m.selectedDifference

	if !changed {
		return cmds
	}

	m.selectedDifference = &selection

	if selection.AfterBlob != nil {
		id := selection.AfterBlob.BlobId
		if m.blobsCache[*id] == nil {
			cmds = append(cmds, m.context.getBlobCmd(id))
		}
	}

	if selection.BeforeBlob != nil {
		id := selection.BeforeBlob.BlobId
		if m.blobsCache[*id] == nil {
			cmds = append(cmds, m.context.getBlobCmd(id))
		}
	}

	m.updateViewport()

	return cmds
}

func (m *model) updateViewport() {
	var before, after string
	if m.selectedDifference.BeforeBlob != nil {
		blob := m.blobsCache[*m.selectedDifference.BeforeBlob.BlobId]
		if blob == nil {
			m.viewport.SetContent("")
			return
		}
		before = string(blob)
	}
	if m.selectedDifference.AfterBlob != nil {
		blob := m.blobsCache[*m.selectedDifference.AfterBlob.BlobId]
		if blob == nil {
			m.viewport.SetContent("")
			return
		}
		after = string(blob)
	}
	edits := myers.ComputeEdits("", before, after)
	diff := gotextdiff.ToUnified("before", "after", before, edits)
	m.viewport.SetContent(fmt.Sprint(diff))
}

func (m model) isBlobOfSelectedDifference(id string) bool {
	return m.selectedDifference.AfterBlob != nil &&
		*m.selectedDifference.AfterBlob.BlobId == id ||
		m.selectedDifference.BeforeBlob != nil &&
			*m.selectedDifference.BeforeBlob.BlobId == id
}
