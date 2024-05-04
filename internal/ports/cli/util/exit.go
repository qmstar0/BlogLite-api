package util

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var ExitModel = exitModel{msg: lipgloss.NewStyle().
	Foreground(lipgloss.Color("239")).
	PaddingLeft(2).Render("Press any key to exit"),
}

type exitModel struct {
	msg string
}

func (m exitModel) Init() tea.Cmd {
	return nil
}

func (m exitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m exitModel) View() string {
	return m.msg
}
