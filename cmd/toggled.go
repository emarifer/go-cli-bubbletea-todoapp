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
	"log"
	"os"
	"strconv"

	"github.com/emarifer/go-cli-bubbletea-todoapp/internal/models"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// newToggledCommand return a cobra.Command pointer that
// represents the `detail` command
func newToggledCommand(db *gorm.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "toggled [id<integer>]",
		Short: "Command to mark a task as completed/incomplete",
		Long: `
This command toggles the status of a task between
completed/incomplete by giving it the task ID (integer).`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Println(err)

				os.Exit(1)
			}

			tasksEdited := models.GetByID(db, id)
			if tasksEdited.ID == 0 {
				fmt.Printf(
					Maroon+Bold+
						"No task could be found with the ID #%d\n"+
						Reset,
					id,
				)

				return
			}

			tasksEdited.Completed = !tasksEdited.Completed

			task := models.UpdateByID(db, id, tasksEdited)

			fmt.Printf(Lime+Bold+"\n\nTask with ID #%d has been changed status\n"+Reset, task.ID)
		},
	}
}

/*
// toggledCmd represents the toggled command
var toggledCmd = &cobra.Command{
	Use:   "toggled [id<integer>]",
	Short: "Command to mark a task as completed/incomplete",
	Long: `
This command toggles the status of a task between
completed/incomplete by giving it the task ID (integer).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toggled called")
	},
}

func init() {
	rootCmd.AddCommand(toggledCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggledCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggledCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
