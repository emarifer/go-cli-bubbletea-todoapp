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

// newDeleteCommand return a cobra.Command pointer that
// represents the `delete` command
func newDeleteCommand(db *gorm.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "delete [id<integer>]",
		Short: "Command to delete a task",
		Long: `
This command deletes a task given its ID (integer).
It returns a warning message to the user if no task with that ID exists.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Println(err)

				os.Exit(1)
			}

			if models.DeleteByID(db, id).RowsAffected == 0 {
				fmt.Printf(
					Maroon+Bold+
						"Could not find/delete any task with ID #%d\n"+
						Reset,
					id,
				)

				return
			}

			fmt.Printf(Lime+Bold+"\n\nTask with ID #%d has been successfully deleted\n"+Reset, id)
		},
	}
}

/*
// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id<integer>]",
	Short: "Command to delete a task",
	Long: `
This command deletes a task given its ID (integer).
It throws a warning message to the user if no task with that ID exists.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
