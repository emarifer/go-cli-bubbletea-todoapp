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
package cmd

import (
	"fmt"

	"github.com/emarifer/go-cli-bubbletea-todoapp/internal/models"
	"github.com/emarifer/go-cli-bubbletea-todoapp/ui/table"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	uitable "github.com/evertras/bubble-table/table"
)

// newListCommand return a cobra.Command pointer that
// represents the `list` command
func newListCommand(db *gorm.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "list",
		Short: "Command to list all tasks",
		Long: `
This command displays a nice table of all the task list.`,
		Run: func(cmd *cobra.Command, args []string) {
			tasksList := models.GetAll(db)
			if len(tasksList) == 0 {
				fmt.Println(Orange + Bold + "There are no tasks to do ¯\\_(ツ)_/¯" + Reset)
				return
			}

			values := []uitable.Row{}
			for _, task := range tasksList {
				done := "❌" // "✅"
				incompleteStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#af00d7"))
				completeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00d700")).Bold(true)
				var row uitable.Row

				if task.Completed {
					done = "✅"
					row = uitable.NewRow(uitable.RowData{
						table.ColumnKeyID:          task.ID,
						table.ColumnKeyName:        task.Name,
						table.ColumnKeyDescription: task.Description,
						table.ColumnKeyStatus:      done,
					}).WithStyle(completeStyle)
				} else {
					row = uitable.NewRow(uitable.RowData{
						table.ColumnKeyID:          task.ID,
						table.ColumnKeyName:        task.Name,
						table.ColumnKeyDescription: task.Description,
						table.ColumnKeyStatus:      done,
					}).WithStyle(incompleteStyle)
				}

				values = append(values, row)
			}

			table.NewModel(values)
		},
	}
}

/*
// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Command to list all tasks",
	Long: `
This command displays a nice table of all the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
