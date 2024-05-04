package cmds

import (
	"context"
	"errors"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/ports/cli/mark"
	"go-blog-ddd/internal/ports/cli/util"
	"time"
)

type EnterCallback interface {
	Enter(app *application.App, item query.CategoryView) (EnterCallback, tea.Cmd)
	View() string
}

type CategoryList struct {
	*util.BaseModel

	items      []query.CategoryView
	totalItems int
	cursor     int
	paginator  paginator.Model

	view func() string
}

func NewCategoryList(app *application.App) *CategoryList {
	p := paginator.New()
	p.Type = paginator.Arabic

	return &CategoryList{
		BaseModel: util.NewBaseModel(app, "Select the category to operate on"),
		paginator: p,
	}
}

func (m *CategoryList) Init() tea.Cmd {
	m.view = func() string { return "" }
	m.items = nil
	m.totalItems = 0
	m.cursor = 0
	return func() tea.Msg {
		ctx, cc := context.WithTimeout(context.Background(), time.Second*5)
		defer cc()

		categorys, err := m.App.Queries.Categorys.GetCategorys(ctx)
		if err != nil {
			return util.CommandInitError{Err: err}
		}

		if categorys.Count <= 0 {
			return util.CommandInitError{Err: errors.New("the total number of data found is 0")}
		}

		m.items = categorys.Items
		m.paginator.SetTotalPages(len(m.items))
		return util.CommandInitSuccess{}
	}
}

func (m *CategoryList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m.CtrlC()
		case "alt+left":
			return m.AltLeft()

		case "enter":
			var cmd tea.Cmd
			m.Next, cmd = m.Next.Update(util.SelectedItem[query.CategoryView]{
				Item: m.GetSelectedItem(),
			})
			return m.Next, tea.Sequence(cmd, m.Next.Init())
		case "left":
			m.paginator.PrevPage()
		case "right":
			m.paginator.NextPage()
		}

	case util.CommandInitSuccess:
		m.view = func() string {
			view, _ := util.CategoryView(m.items[m.paginator.Page])
			return lipgloss.JoinVertical(
				lipgloss.Left,
				util.Title(m.Title, m.paginator.View()),
				util.BlockLeftLineStyle.Render(view),
			)
		}
		return m, nil

	case util.CommandInitError:
		return m.Prev, tea.Sequence(
			util.Record("Failed to load data", mark.Error, msg.Err.Error()),
			m.Prev.Init(),
		)

	case util.SubmitSuccess, util.SubmitError:
		return m.Next, nil
	}
	return m, nil
}

func (m *CategoryList) View() string {
	return m.view()
}

func (m *CategoryList) GetSelectedItem() query.CategoryView {
	return m.items[m.paginator.Page]
}
