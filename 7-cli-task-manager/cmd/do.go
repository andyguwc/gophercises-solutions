
package cmd

import (
	"fmt"
	"strconv"
	// "os"

	"github.com/spf13/cobra"
	"github.com/andyguwc/go-course/gophercises/7-cli-task-manager/db"

)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("not parsing arg", arg)
			} else {
				ids = append(ids, id)
			}
		}

		// tasks, err := db.AllTasks()
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	os.Exit(1)
		// }

		for _, id := range ids {
			db.DeleteTask(id)
		}


		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
