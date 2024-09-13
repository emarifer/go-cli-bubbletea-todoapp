/*
Copyright © 2024 Enrique Marín

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package table

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	uitable "github.com/evertras/bubble-table/table"
)

const (
	ColumnKeyID          = "id"
	ColumnKeyName        = "name"
	ColumnKeyDescription = "description"
	ColumnKeyStatus      = "status"

	Bold   = "\033[1m"
	Reset  = "\033[0m"
	White  = "\033[97m"
	Yellow = "\033[33m"
)

var (
	customBorder = uitable.Border{
		Top:    "─",
		Left:   "│",
		Right:  "│",
		Bottom: "─",

		TopRight:    "╮",
		TopLeft:     "╭",
		BottomRight: "╯",
		BottomLeft:  "╰",

		TopJunction:    "╥",
		LeftJunction:   "├",
		RightJunction:  "┤",
		BottomJunction: "╨",
		InnerJunction:  "╫",

		InnerDivider: "║",
	}
)

type Model struct {
	tableModel uitable.Model
}

func NewModel(rows []uitable.Row) {
	columns := []uitable.Column{
		uitable.NewColumn(ColumnKeyID, "ID", 5).WithStyle(
			lipgloss.NewStyle().
				Faint(true).
				Foreground(lipgloss.Color("#88f")).
				Align(lipgloss.Center)),
		uitable.NewColumn(ColumnKeyName, "Name", 21),
		uitable.NewColumn(ColumnKeyDescription, "Description", 30),
		uitable.NewColumn(ColumnKeyStatus, "Status", 8).WithStyle(
			lipgloss.NewStyle().Align(lipgloss.Center)),
	}

	// Start with the default key map and change it slightly, just for demoing
	keys := uitable.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down", "s")
	keys.RowUp.SetKeys("k", "up", "w")

	model := Model{
		// Throw features in... the point is not to look good, it's just reference!
		tableModel: uitable.New(columns).
			WithRows(rows).
			HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)).
			SelectableRows(true).
			Focused(true).
			Border(customBorder).
			WithKeyMap(keys).
			WithStaticFooter("Footer!").
			WithPageSize(10).
			WithSelectedText(" ", "✓").
			WithBaseStyle(
				lipgloss.NewStyle().
					BorderForeground(lipgloss.Color("#a38")).
					Foreground(lipgloss.Color("#a7a")).
					Align(lipgloss.Left),
			).
			// SortByAsc(ColumnKeyID).
			WithMissingDataIndicatorStyled(uitable.StyledCell{
				Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#faa")),
				Data:  "<nil>",
			}),
	}

	model.updateFooter()

	// return model
	if _, err := tea.NewProgram(model).Run(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) updateFooter() {
	highlightedRow := m.tableModel.HighlightedRow()
	rows := m.tableModel.GetVisibleRows()
	pending := 0
	for _, row := range rows {
		if row.Data["status"] == "❌" {
			pending++
		}
	}

	footerText := fmt.Sprintf(
		`• Pg. `+White+`%d/%d`+Reset+`
• Pending tasks: `+White+`%d`+Reset+`

📄 `+Yellow+Bold+`Description of the current task:`+Reset+`
	%s`,
		m.tableModel.CurrentPage(),
		m.tableModel.MaxPages(),
		pending,
		highlightedRow.Data[ColumnKeyDescription],
	)

	m.tableModel = m.tableModel.WithStaticFooter(footerText)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.tableModel, cmd = m.tableModel.Update(msg)
	cmds = append(cmds, cmd)

	// We control the footer text, so make sure to update it
	m.updateFooter()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)

		case "i":
			m.tableModel = m.tableModel.WithHeaderVisibility(!m.tableModel.GetHeaderVisibility())

		case "f":
			m.tableModel = m.tableModel.WithFooterVisibility(!m.tableModel.GetFooterVisibility())
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	body := strings.Builder{}

	body.WriteString("\nPress down/up or 'j'/'k' to move through the rows\n")
	body.WriteString("Press left/right or PgUp/PgDown to move pages\n")
	body.WriteString("Press 'i' to toggle the header visibility\n")
	body.WriteString("Press 'f' to toggle the footer visibility\n")
	body.WriteString("Press space/enter to select a row, q or ctrl+c to quit\n\n")

	body.WriteString(m.tableModel.View())

	body.WriteString("\n")

	return body.String()
}

/* REFERENCES:
https://github.com/Evertras/bubble-table/blob/main/examples/features/main.go
*/
