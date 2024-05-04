package util

import "github.com/charmbracelet/lipgloss"

var (
	UnderlineStyle     = lipgloss.NewStyle().Underline(true)
	BlodStyle          = lipgloss.NewStyle().Bold(true)
	BlodUnderlineStyle = lipgloss.NewStyle().Underline(true).Bold(true)
)

var (
	ItemOffset  = lipgloss.NewStyle().PaddingLeft(2)
	NormalTitle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
			Padding(0, 0, 0, 2)

	NormalDesc = NormalTitle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	SelectedTitle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
			Padding(0, 0, 0, 1)

	SelectedDesc = SelectedTitle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})
)

var (
	DocStyle           = lipgloss.NewStyle().Padding(1, 4)
	BlockLeftLineStyle = lipgloss.NewStyle().PaddingLeft(1).Border(lipgloss.NormalBorder(), false, false, false, true)
)
