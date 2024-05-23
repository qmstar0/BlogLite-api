package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/apps"
	"go-blog-ddd/internal/ports/cli/util"
)

type Cli struct {
	app *apps.DomainApp
}

func NewCli(app *apps.DomainApp) *Cli {
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
