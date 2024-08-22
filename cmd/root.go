package cmd

import (
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "a simple task cli for you",
	Long: `a simple task cli you can use to manage your daily tasks,
you can add, delete and update your tasks,
you can mark them as done or in-progress or leave it as todo,
you can also see the list of done, todo or in progress tasks.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(statics.ErrorStyle.Render(err.Error()))
		os.Exit(1)
	}
}
