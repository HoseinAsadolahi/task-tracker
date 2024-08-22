package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/pkg/repository"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"github.com/spf13/cobra"
	"strconv"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task description with specified id",
	Long:  `simply updates description of a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("update command expects 2"+
				" arguments, got %d", len(args))))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(statics.ErrorStyle.Render("update command expects integer as id"))
			return
		}
		repository.UpdateTask(id, args[1], "")
	},
}

var todoCmd = &cobra.Command{
	Use:   "mark-todo",
	Short: "update task status to \"todo\" with specified id",
	Long:  `simply updates status of a task to "todo"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("mark-todo command expects 1"+
				" arguments, got %d", len(args))))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(statics.ErrorStyle.Render("mark-todo command expects integer as id"))
			return
		}
		repository.UpdateTask(id, "", "todo")
	},
}

var doneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "update task status to \"done\" with specified id",
	Long:  `simply updates status of a task to "done"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("mark-done command expects 1"+
				" arguments, got %d", len(args))))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(statics.ErrorStyle.Render("mark-done command expects integer as id"))
			return
		}
		repository.UpdateTask(id, "", "done")
	},
}

var inProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "update task status to \"in-progress\" with specified id",
	Long:  `simply updates status of a task to "in-progress"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(statics.ErrorStyle.Render(fmt.Sprintf("mark-in-progress command expects 1"+
				" arguments, got %d", len(args))))
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(statics.ErrorStyle.Render("mark-in-progress command expects integer as id"))
			return
		}
		repository.UpdateTask(id, "", "in-progress")
	},
}
