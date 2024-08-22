package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/pkg/repository"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task with specified id",
	Long:  `simply deletes a task from current tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("delete command expects 1"+
				" arguments, got %d", len(args))))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(statics.ErrorStyle.Render("delete command expects integer as argument"))
			return
		}
		repository.DeleteTask(id)
	},
}
