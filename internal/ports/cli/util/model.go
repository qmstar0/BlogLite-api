package util

import (
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/cli/mark"
)

type BaseModel struct {
	App *application.App

	Title string

	Prev tea.Model
	Next tea.Model
}

func NewBaseModel(app *application.App, title string) *BaseModel {
	return &BaseModel{
		App:   app,
		Title: title,
		Prev:  nil,
		Next:  ExitModel,
	}
}

func (m *BaseModel) With(prev, next tea.Model) {
	m.Prev = prev
	m.Next = next
}

func (m *BaseModel) CtrlC() (tea.Model, tea.Cmd) {
	return ExitModel, tea.Sequence(
		Record(m.Title, mark.Error, "Cancel operation"),
		tea.Quit,
	)
}

func (m *BaseModel) AltLeft() (tea.Model, tea.Cmd) {
	return m.Prev, tea.Sequence(
		Record(m.Title, mark.Error, "Cancel operation"),
		m.Prev.Init(),
	)
}
