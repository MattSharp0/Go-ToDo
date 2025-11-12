package internal

import (
	"todo-list-cli/internal/terminal"
)

func CalculateDisplayColumns(colPercentages ...int) (int, []int) {
	columns := len(colPercentages)
	output := make([]int, columns)

	var totalPercent int
	for _, pct := range colPercentages {
		totalPercent += pct
	}

	w, _ := terminal.GetTerminalSize()

	// Calculate widths from percentages
	var totalWidth int
	for i := 0; i < columns; i++ {
		output[i] = (colPercentages[i] * w) / 100 // Fixed: percentage * width / 100
		totalWidth += output[i]
	}

	// If percentages don't add up to 100, adjust the last column
	if totalPercent < 100 {
		remainingWidth := w - totalWidth
		output[columns-1] += remainingWidth
	} else if totalPercent > 100 {
		// If percentages exceed 100, scale them down proportionally
		scale := w
		for i := 0; i < columns; i++ {
			output[i] = (colPercentages[i] * scale) / totalPercent
		}
	}

	return w, output
}
