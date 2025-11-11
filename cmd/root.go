package cmd

import (
	"fmt"
	"os"
	"todo-list-cli/internal"

	"github.com/spf13/cobra"
)

var ToDoList *internal.ToDoList

var rootCmd = &cobra.Command{
	Use:   "GoToDo",
	Short: "Go ToDo is a CLI based ToDo list",
	Long:  "A fast, lightweight ToDo list app that lives in the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Go-ToDo, a CLI based ToDo list application written in Go.")
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
