
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/andyguwc/go-course/gophercises/7-cli-task-manager/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks ",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0{
			fmt.Println("no tasks")
			return 
		}
		fmt.Println("you have the following tasks")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
