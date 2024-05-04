package util

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Countinue struct {
	*BaseModel
	msg string
}

func NewCountinue(msg string) Countinue {
	return Countinue{
		BaseModel: NewBaseModel(nil, msg),
		msg:       msg,
	}
}

func (m Countinue) Init() tea.Cmd {
	return nil
}

func (m Countinue) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m.Next, m.Next.Init()
	}
	return m, nil
}

func (m Countinue) View() string {
	return m.msg
}
