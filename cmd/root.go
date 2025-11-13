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
