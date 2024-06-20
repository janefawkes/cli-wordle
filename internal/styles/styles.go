package styles

import "github.com/charmbracelet/lipgloss"

var (
	CellStyle  = lipgloss.NewStyle().Width(3).Height(1).Align(lipgloss.Center).Padding(0, 1).Margin(0, 0)
	Absent     = CellStyle.Background(lipgloss.Color("#787C7E")).Foreground(lipgloss.Color("#FFFFFF"))
	Present    = CellStyle.Background(lipgloss.Color("#C9B458")).Foreground(lipgloss.Color("#FFFFFF"))
	Correct    = CellStyle.Background(lipgloss.Color("#6AAA64")).Foreground(lipgloss.Color("#FFFFFF"))
	DefaultBox = CellStyle.Background(lipgloss.Color("#FFFFFF")).Foreground(lipgloss.Color("#000000"))
)
