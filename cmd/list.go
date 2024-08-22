package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/pkg/repository"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks or tasks with specified status",
	Long:  `simply lists all tasks or tasks with specified status`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			repository.ListAllTasks()
			return
		}
		if len(args) > 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("list command take at most 1 argument,"+
				" got %d", len(args))))
			return
		}
		if args[0] != "todo" && args[0] != "in-progress" && args[0] != "done" {
			fmt.Println(statics.ErrorStyle.Render("the argument should be done, in-progress or todo!"))
			return
		}
		repository.ListTasksByStatus(args[0])
	},
}
