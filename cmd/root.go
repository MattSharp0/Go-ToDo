package cmd

import (
	"fmt"
	"os"
	"todo/internal"

	"github.com/spf13/cobra"
)

var ToDoList *internal.ToDoList

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "ToDo is a CLI based ToDo list written in Go",
	Long:  "A simple, lightweight ToDo list app that lives in the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ToDo is a CLI based ToDo list written in Go.\nRun 'todo list' to show todos,'todo add <item>' to add a todo or 'todo help' for more infomration")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func SetToDoList(tdl *internal.ToDoList) {
	ToDoList = tdl
}
