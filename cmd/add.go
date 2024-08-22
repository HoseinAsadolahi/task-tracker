package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/pkg/repository"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task with the specified description",
	Long:  `simply adds a task to current tasks and set the status to "todo"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("add command expects 1 arguments, got %d", len(args))))
			return
		}
		repository.AddTask(args[0])
	},
}
