package cmd

import (
	"todo/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listToDosCommand)

	listToDosCommand.Flags().BoolP("incomplete", "i", false, "Show only incomplete todos")
	listToDosCommand.Flags().BoolP("complete", "c", false, "Show only complete todos")
	listToDosCommand.Flags().BoolP("today", "t", false, "Show todos created today")
	listToDosCommand.Flags().BoolP("this-week", "w", false, "Show todos created this week")
	listToDosCommand.Flags().BoolP("this-month", "m", false, "Show todos created this month")
}

var listToDosCommand = &cobra.Command{
	Use:   "list",
	Short: "List all ToDo items",
	Long:  "List all ToDo items in a tabular format",
	Run: func(cmd *cobra.Command, args []string) {
		incomplete, _ := cmd.Flags().GetBool("incomplete")
		complete, _ := cmd.Flags().GetBool("complete")
		today, _ := cmd.Flags().GetBool("today")
		thisWeek, _ := cmd.Flags().GetBool("this-week")
		thisMonth, _ := cmd.Flags().GetBool("this-month")

		// Build filter
		filter := &internal.ToDoFilter{
			Complete:        internal.FilterAll,
			CreatedInPeriod: internal.PeriodAll,
		}

		if incomplete {
			filter.Complete = internal.FilterIncomplete
		} else if complete {
			filter.Complete = internal.FilterComplete
		}

		if today {
			filter.CreatedInPeriod = internal.PeriodToday
		} else if thisWeek {
			filter.CreatedInPeriod = internal.PeriodThisWeek
		} else if thisMonth {
			filter.CreatedInPeriod = internal.PeriodThisMonth
		}

		ToDoList.DisplayTodoList(filter)

	},
}
