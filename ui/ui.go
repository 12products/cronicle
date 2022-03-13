package ui

import (
	"cronicle/ui/components/help"
	"cronicle/ui/components/tabs"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"

	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	currSection int
	tabs        tabs.Model
	keys        utils.KeyMap
	help        help.Model
	ctx         context.Context
}

func New() Model {
	return Model{
		currSection: 0,
		tabs:        tabs.New(),
		keys:        utils.Keys,
		help:        help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd     tea.Cmd
		cmds    []tea.Cmd
		helpCmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevSection):
			prevSection := m.getPrevSection()
			m.setCurrSection(prevSection)

		case key.Matches(msg, m.keys.NextSection):
			nextSection := m.getNextSection()
			m.setCurrSection(nextSection)

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)
	}

	m.help, helpCmd = m.help.Update(msg)
	cmds = append(cmds, cmd, helpCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.tabs.View(m.ctx))

	s.WriteString("\n")

	s.WriteString(m.help.View(m.ctx))

	return s.String()
}

func (m *Model) setCurrSection(newSection int) {
	m.currSection = newSection
	m.tabs.SetCurrSection(newSection)
}

func (m Model) getNextSection() int {
	return (m.currSection + 1) % len(constants.Sections)
}

func (m Model) getPrevSection() int {
	m.currSection = (m.currSection - 1) % len(constants.Sections)

	if m.currSection < 0 {
		m.currSection += len(constants.Sections)
	}

	return m.currSection
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.help.SetWidth(msg.Width)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
}
