package cmds

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/ports/cli/mark"
	"go-blog-ddd/internal/ports/cli/util"
	"time"
)

type DeleteCategory struct {
	*util.BaseModel
	ask  bool
	next tea.Model

	Item query.CategoryView
}

func NewDeleteCategoryModel(app *application.App) *DeleteCategory {
	return &DeleteCategory{
		BaseModel: util.NewBaseModel(app, "Delete category"),
		ask:       false,
		Item:      query.CategoryView{},
	}
}

func (m *DeleteCategory) Init() tea.Cmd {
	m.ask = false
	return nil
}

func (m *DeleteCategory) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m.CtrlC()
		case "alt+left":
			return m.AltLeft()

		case "left":
			m.ask = true

		case "right":
			m.ask = false

		case "enter":
			if m.ask {
				return m, m.Submit()
			}
			return m.Prev, nil
		}
	case util.SelectedItem[query.CategoryView]:
		m.Item = msg.Item
		return m, nil

	case util.SubmitSuccess:
		return m.Next, util.Record(m.Title, mark.OKNotice, util.BlodUnderlineStyle.Render(m.Item.Name))

	case util.SubmitError:
		return m.Next, util.Record(m.Title, mark.Error, fmt.Sprintf("error: %s", msg.Err.Error()))
	}
	return m, nil
}

func (m *DeleteCategory) View() string {
	var s = fmt.Sprintf("yes/%s", util.BlodUnderlineStyle.Render("No"))
	if m.ask {
		s = fmt.Sprintf("%s/No", util.BlodUnderlineStyle.Render("yes"))
	}

	return fmt.Sprintf("%s Confirm again to delete '%s'? %s %s",
		mark.Progressing,
		util.BlodUnderlineStyle.Render(m.Item.Name),
		mark.Right,
		s,
	)
}

func (m *DeleteCategory) Submit() tea.Cmd {
	c, cc := context.WithTimeout(context.Background(), time.Second*5)
	defer cc()

	err := m.App.Commands.DeleteCategory.Handle(c, commands.DeleteCategory{ID: m.Item.ID})
	if err != nil {
		return func() tea.Msg { return util.SubmitError{Err: err} }
	}
	return func() tea.Msg { return util.SubmitSuccess{} }
}
