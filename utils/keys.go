package utils

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	NextSection key.Binding
	PrevSection key.Binding
	Help        key.Binding
	Quit        key.Binding
}

var (
	Keys = KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		PrevSection: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "previous section"),
		),
		NextSection: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "next section"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
)