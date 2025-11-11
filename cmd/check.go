package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkToDoCommand)

	checkToDoCommand.Flags().BoolP("complete", "c", false, "Mark as Complete")
	checkToDoCommand.Flags().BoolP("incomplete", "i", false, "Mark as Incomplete")

}

var checkToDoCommand = &cobra.Command{
	Use:   "check [itemId]",
	Short: "Check/Uncheck ToDo Item",
	Long:  "Check/Uncheck a ToDo list item by ID",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		incomplete, _ := cmd.Flags().GetBool("incomplete")

		complete := true
		if incomplete {
			complete = false
		}

		itemIdArg := args[0]
		itemId, err := strconv.Atoi(itemIdArg)
		if err != nil {
			fmt.Printf("Unable to parse ItemID: %v", itemIdArg)
			return
		}

		ToDoList.CheckToDo(itemId, complete)

	},
}
