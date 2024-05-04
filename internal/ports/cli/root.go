package cli

import (
	"fmt"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/cli/mark"
	"go-blog-ddd/internal/ports/cli/util"
)

type MainModel struct {
	*util.BaseModel

	Items      []Item
	TotalItems int

	index     int
	paginator paginator.Model
}

func InitRootModel(app *application.App) *MainModel {

	p := paginator.New()

	p.Type = paginator.Dots
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	model := &MainModel{
		BaseModel: util.NewBaseModel(app, "What do you want to do"),
		index:     0,
		paginator: p,
	}
	model.With(util.ExitModel, nil)
	model.SetItems(InitItems(app, model))

	return model
}

func (m *MainModel) Init() tea.Cmd {
	m.index = 0
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m.Prev, tea.Sequence(
				util.Record(m.Title, mark.Point, "Bye!"),
				tea.Quit,
			)
		case "up":
			m.CursorUp()
		case "down":
			m.CursorDown()
		case "left":
			m.paginator.PrevPage()
		case "right":
			m.paginator.NextPage()
			if m.paginator.OnLastPage() {
				itemOnPage := m.paginator.ItemsOnPage(m.TotalItems)
				if m.index >= itemOnPage {
					m.index = itemOnPage - 1
				}
			}
		case "enter":
			item := m.GetSelectedItem()
			return item.Model, tea.Sequence(
				util.Record("You chose", mark.Point, item.Title),
				item.Model.Init(),
			)
		}
	case util.Hallo:
		return m, util.Record("Welcome back", mark.Point, "Admin")
	}
	return m, nil
}

func (m *MainModel) View() string {
	start, end := m.paginator.GetSliceBounds(m.TotalItems)

	views := make([]string, end-start)
	for i, item := range m.Items[start:end] {
		if i == m.index {
			views[i] = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7486FB")).
				Render(fmt.Sprintf("→ %s", item.Title))
		} else {
			views[i] = lipgloss.NewStyle().PaddingLeft(2).Render(item.Title)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		util.Title(m.Title, m.paginator.View()),
		util.BlockLeftLineStyle.Render(lipgloss.JoinVertical(lipgloss.Left, views...)),
	)
}

func (m *MainModel) CursorUp() {
	m.index--
	if m.index < 0 && m.paginator.Page == 0 {
		m.index = 0
		return
	}
	if m.index >= 0 {
		return
	}
	m.paginator.PrevPage()
	m.index = m.paginator.PerPage - 1
}

func (m *MainModel) CursorDown() {
	itemOnPage := m.paginator.ItemsOnPage(m.TotalItems)
	m.index++
	if m.index < itemOnPage {
		return
	}
	if !m.paginator.OnLastPage() {
		m.paginator.NextPage()
		m.index = 0
		return
	}
	m.index = itemOnPage - 1
}
func (m *MainModel) GetSelectedItem() Item {
	return m.Items[m.paginator.PerPage*m.paginator.Page+m.index]
}

func (m *MainModel) SetItems(items []Item) {
	m.TotalItems = len(items)
	m.Items = items
	m.paginator.PerPage = 7
	m.paginator.SetTotalPages(m.TotalItems)
}
