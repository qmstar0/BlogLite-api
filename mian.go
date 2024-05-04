package main

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/cli"
	"os"
	"strings"
)

type Model struct {
	task []string
}

type M struct {
	task []string
}

func (m Model) Init() tea.Cmd {
	fmt.Println("model")
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch s := msg.String(); s {
		case "q":
			return M{}, nil
		default:
			m.task = append(m.task, s)
			return m, nil
		}
	case error:
		m.task = append(m.task, "find err")
		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("%s", strings.Join(m.task, ";"))
}

func (m M) Init() tea.Cmd {
	return tea.Sequence(
		tea.Printf("不可以直接从M开始"),
		tea.Quit,
	)
}

func (m M) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "p":
			return Model{task: make([]string, 0)}, func() tea.Msg {
				return errors.New("test")
			}
		default:

			m.task = append(m.task, msg.String())
			return m, nil
		}
	}
	return m, nil
}

func (m M) View() string {
	return "M" + strings.Join(m.task, ": ")
}

func main() {
	//_, err := tea.NewProgram(M{task: make([]string, 0)}).Run()
	//if err != nil {
	//	panic(err)
	//}
	//_, err := tea.NewProgram(domain.InitRootModel(application.NewApp()), tea.WithAltScreen()).Run()
	//if err != nil {
	//	panic(err)
	//}
	err := cli.NewCli(application.NewApp()).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
