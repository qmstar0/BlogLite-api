package mark

import "github.com/charmbracelet/lipgloss"

var markStyle = lipgloss.NewStyle()

var (
	Point = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("•")
	Error = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("×")

	OK          = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("✓")
	OKNotice    = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("✓")
	Notice      = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("!")
	Progressing = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("?")

	Down        = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("↓")
	Up          = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("↑")
	Arrow       = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("→")
	Left        = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("‹")
	Right       = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("›")
	DoubleLeft  = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("«")
	DoubleRight = markStyle.Foreground(lipgloss.Color("#3BF3FF")).Render("»")
)
