package util

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"go-blog-ddd/internal/ports/cli/mark"
)

type SubmitSuccess struct{}

type SubmitError struct{ Err error }

type CommandInitSuccess struct {
}

type CommandInitError struct {
	Err error
}

type Hallo struct {
	Msg string
}
type SelectedItem[I any] struct {
	Item I
}

func Record(msg, state, value string) tea.Cmd {
	return tea.Printf("%s %s %s %s",
		state, msg, mark.DoubleRight, value,
	)
}

func Title(msg, value string) string {
	return fmt.Sprintf("%s %s %s %s", mark.Right, msg, mark.Right, value)
}
