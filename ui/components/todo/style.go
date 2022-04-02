package todo

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	cellStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			MaxHeight(1)

	selectedCellStyle = cellStyle.Copy().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#FFFFFF"))

	titleCellStyle = cellStyle.Copy().
			Bold(true)

	// singleRuneTitleCellStyle = titleCellStyle.Copy().Width(1)

	rowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true)
)
