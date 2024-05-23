package cmds

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"go-blog-ddd/internal/apps"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/ports/cli/mark"
	"go-blog-ddd/internal/ports/cli/util"
	"strings"
	"time"
)

type CreateCategory struct {
	*util.BaseModel

	index int

	label            [3]string
	inputPlaceholder [2]string

	input textinput.Model
	Item  commands.CreateCategory
}

func NewCreateCategory(app *apps.DomainApp) *CreateCategory {
	return &CreateCategory{
		BaseModel: util.NewBaseModel(app, "Create category"),
		index:     0,
		input:     textinput.New(),
		Item:      commands.CreateCategory{},
		label: [3]string{
			fmt.Sprintf("Please enter a category name %s ", mark.Right),
			fmt.Sprintf("Description of the category %s ", mark.Right),
			fmt.Sprintf("Please enter a category name (cannot be null) %s ", mark.Right),
		},
	}
}

func (m *CreateCategory) Init() tea.Cmd {
	m.index = 0
	m.Item = commands.CreateCategory{}
	m.input.Placeholder = m.inputPlaceholder[m.index]
	m.input.Prompt = m.label[m.index]
	m.input.SetValue("")
	return m.input.Focus()
}

func (m *CreateCategory) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m.CtrlC()
		case "alt+left":
			return m.AltLeft()

		case "enter":
			switch m.index {
			case 0:
				m.Item.Name = strings.TrimSpace(m.input.Value())

				//检验
				if m.Item.Name == "" {
					m.input.Prompt = m.label[2]
					return m, nil
				}

				m.index++
				m.input.Placeholder = m.inputPlaceholder[m.index]
				m.input.Prompt = m.label[m.index]
				m.input.SetValue("")
				return m, nil
			case 1:
				m.input.Blur()
				m.index++
				m.Item.Desc = strings.TrimSpace(m.input.Value())
				return m, m.Submit()
			case 2:
				return m, nil
			}
		}

	case util.SubmitSuccess:
		return m.Next, util.Record(m.Title, mark.OK, util.BlodUnderlineStyle.Render(m.Item.Name))

	case util.SubmitError:
		return m.Next, util.Record(m.Title, mark.Error, fmt.Sprintf("error: %s", msg.Err.Error()))
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
func (m *CreateCategory) View() string {
	var section = make([]string, 0, 2)
	if m.Item.Name != "" {
		section = append(section, fmt.Sprintf("%s %s%s",
			mark.OK, m.label[0], m.Item.Name))
	}

	if m.index < 2 {
		section = append(section,
			fmt.Sprintf("%s %s", mark.Progressing, m.input.View()),
		)
	} else {
		section = append(section, fmt.Sprintf("%s %s%s",
			mark.OK,
			m.input.Prompt,
			m.input.Value()))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		util.Title(m.Title, "…"),
		util.BlockLeftLineStyle.Render(lipgloss.JoinVertical(
			lipgloss.Left,
			section...,
		)),
	)
}

func (m *CreateCategory) Submit() tea.Cmd {
	return func() tea.Msg {
		timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
		defer cancelFunc()

		_, err := m.App.Commands.CreateCategory.Handle(timeout, m.Item)

		if err != nil {
			return util.SubmitError{Err: err}
		}
		return util.SubmitSuccess{}
	}
}
