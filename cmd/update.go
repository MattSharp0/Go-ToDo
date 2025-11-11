package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateToDoCommand)

}

var updateToDoCommand = &cobra.Command{
	Use:   "update [ToDo Id] [ItemName]",
	Short: "Update ToDo Item",
	Long:  "Update ToDo Item text",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		itemIdArg := args[0]
		itemId, err := strconv.Atoi(itemIdArg)
		if err != nil {
			fmt.Printf("Unable to parse ItemID: %v", itemIdArg)
			return
		}

		newItemTextArgs := args[1:]
		newItemText := strings.Join(newItemTextArgs, " ")

		ToDoList.UpdateToDoItem(itemId, newItemText)

	},
}
