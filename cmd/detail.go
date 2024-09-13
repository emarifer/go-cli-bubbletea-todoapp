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
	"time"

	"github.com/emarifer/go-cli-bubbletea-todoapp/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// newDetailCommand return a cobra.Command pointer that
// represents the `detail` command
func newDetailCommand(db *gorm.DB) *cobra.Command {

	return &cobra.Command{
		Use:   "detail [id<integer>]",
		Short: "Command to show a task detail",
		Long: `
This command displays the details of the task specified by its ID.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Println(err)

				os.Exit(1)
			}

			task := models.GetByID(db, id)
			if task.ID == 0 {
				fmt.Printf(
					Maroon+Bold+
						"No task could be found with the ID #%d\n"+
						Reset,
					id,
				)

				return
			}

			fmt.Printf(
				Cyan+Bold+"ID: "+Reset+"%d"+
					Cyan+Bold+"\nName: "+Reset+"%s"+
					Cyan+Bold+" \nDescription: "+Reset+"%s"+
					Cyan+Bold+"\nStatus: "+Reset+"%t"+
					Cyan+Bold+" \nCreated At: "+Reset+"%s\n",
				task.ID,
				task.Name,
				task.Description,
				task.Completed,
				task.CreatedAt.Format(time.RFC822Z),
			)
		},
	}
}

/*
// detailCmd represents the detail command
var detailCmd = &cobra.Command{
	Use:   "detail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("detail called")
	},
}

func init() {
	rootCmd.AddCommand(detailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
