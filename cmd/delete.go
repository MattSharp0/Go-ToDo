package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteToDoCommand)

}

var deleteToDoCommand = &cobra.Command{
	Use:   "delete [ToDo Id]",
	Short: "Delete ToDo Item",
	Long:  "Delete ToDo Item",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		itemIdArg := args[0]
		itemId, err := strconv.Atoi(itemIdArg)
		if err != nil {
			fmt.Printf("Unable to parse ItemID: %v", itemIdArg)
			return
		}

		ToDoList.DeleteToDo(itemId)

	},
}
