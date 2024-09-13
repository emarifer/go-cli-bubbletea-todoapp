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
	"strings"

	"github.com/emarifer/go-cli-bubbletea-todoapp/internal/models"
	"github.com/emarifer/go-cli-bubbletea-todoapp/ui/form"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// newAddCommand return a cobra.Command pointer that
// represents the `add` command
func newAddCommand(db *gorm.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "add",
		Short: "Command to add a task",
		Long: `
This command add a task to the task list.`,
		Run: func(cmd *cobra.Command, args []string) {
			task := form.Create()
			n := strings.Trim(task.Name, " ")
			d := strings.Trim(task.Description, " ")
			if len(n) < 2 || len(d) < 2 {
				fmt.Println(Maroon + Bold + "No task has been created!!" + Reset)
				return
			}
			newTask := models.Add(db, task)

			fmt.Printf(
				Lime+Bold+"Task created successfully:"+Reset+" %s [ID #%d]\n"+Reset,
				newTask.Name, newTask.ID,
			)
		},
	}
}

/*
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add a task",
	Long: `
This command add a task to the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		task := form.Create()

		fmt.Printf("Task created successfully: %s(%d)\n", task.Name, task.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
