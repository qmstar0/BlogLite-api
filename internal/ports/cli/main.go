package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/cli/util"
)

type Cli struct {
	app *application.App
}

func NewCli(app *application.App) *Cli {
	return &Cli{app: app}
}

func (c *Cli) Run() error {
	program := tea.NewProgram(InitRootModel(c.app))
	go func() {
		program.Send(util.Hallo{})
	}()
	_, err := program.Run()
	return err
}
