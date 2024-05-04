package cmds

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/ports/cli/mark"
	"go-blog-ddd/internal/ports/cli/util"
	"strings"
	"time"
)

type ModifyCategoryDesc struct {
	*util.BaseModel

	state string

	input textinput.Model
	Item  commands.ModifyCategoryDesc
}

func NewModifyCategroyDesc(app *application.App) *ModifyCategoryDesc {
	t := textinput.New()
	t.Placeholder = "Description"
	return &ModifyCategoryDesc{
		BaseModel: util.NewBaseModel(app, "Modify categroy description"),
		input:     t,
		Item:      commands.ModifyCategoryDesc{},
	}
}

func (m *ModifyCategoryDesc) Init() tea.Cmd {
	m.state = mark.Progressing
	return m.input.Focus()
}

func (m *ModifyCategoryDesc) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m.CtrlC()
		case "alt+left":
			return m.AltLeft()
		case "enter":
			m.state = mark.OK
			m.Item.NewDesc = strings.TrimSpace(m.input.Value())
			return m, m.Submit()
		}
	case util.SelectedItem[query.CategoryView]:
		m.Item = commands.ModifyCategoryDesc{
			ID:      msg.Item.ID,
			NewDesc: msg.Item.Desc,
		}
		m.input.Prompt = fmt.Sprintf("Modify the description of '%s' %s ",
			util.BlodUnderlineStyle.Render(msg.Item.Name),
			mark.Right,
		)
		m.input.SetValue(msg.Item.Desc)
		return m, nil

	case util.SubmitSuccess:
		return m.Next, util.Record(m.Title, mark.OK, "Successfully modified")

	case util.SubmitError:
		return m.Next, util.Record(m.Title, mark.Error, fmt.Sprintf("error: %s", msg.Err.Error()))
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *ModifyCategoryDesc) View() string {
	return util.BlockLeftLineStyle.Render(fmt.Sprintf("%s %s", m.state,
		m.input.View(),
	))
}

func (m *ModifyCategoryDesc) Submit() tea.Cmd {
	return func() tea.Msg {
		timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
		defer cancelFunc()

		err := m.App.Commands.ModifyCategoryDesc.Handle(timeout, m.Item)
		if err != nil {
			return util.SubmitError{Err: err}
		}
		return util.SubmitSuccess{}
	}
}
