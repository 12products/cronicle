package utils

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	NextSection key.Binding
	PrevSection key.Binding
	ScrollUp    key.Binding
	ScrollDown  key.Binding
	Help        key.Binding
	Quit        key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.PrevSection, k.NextSection},
		{k.Help, k.Quit},
	}
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
		ScrollUp: key.NewBinding(
			key.WithKeys("ctrl+u"),
			key.WithHelp("ctrl+u", "scroll document up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("Cctrl+d", "scroll document down"),
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
