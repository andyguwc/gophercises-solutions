package cmd 

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/andyguwc/go-course/gophercises/7-cli-task-manager/db"

)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list ",
	// Long: `xx`,
	Run: func(cmd *cobra.Command, args []string) { 

		task :=strings.Join(args, " ")
		taskId, err := db.CreateTask(task)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added .%d \"%s\" to your list", taskId, task)
	},
}


func init() {
	RootCmd.AddCommand(addCmd)
}