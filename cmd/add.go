package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addToDoCommand)

}

var addToDoCommand = &cobra.Command{
	Use:   "add [ToDo]",
	Short: "Add a ToDo",
	Long:  "Add an Item to the Todo List",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		itemName := strings.Join(args, " ")

		ToDoList.AddToDo(itemName)

	},
}
