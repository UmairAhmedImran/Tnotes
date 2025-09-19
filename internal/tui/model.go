package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	 "github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Screen int


const gap = "\n\n"

const (
	ScreenMain Screen = iota
	ScreenAdd
	ScreenView
	ScreenSearch
)

type Model struct {
	quitting bool
	current  Screen
	subModel tea.Model // holds whichever child model is active
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("41")).Render

func (m Model) helpView() string {
  return helpStyle("            Hello")
}

func New() Model {
	return Model{
		current: ScreenMain,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
    case "esc", "ctrl+c":
			return m, tea.Quit
    }
	}

	// Forward updates to submodel if it exists
	if m.subModel != nil {
		newSub, cmd := m.subModel.Update(msg)
		m.subModel = newSub
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	if m.subModel != nil {
    _, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(2),
	)
  if err != nil {
    return  ""
  }
		return m.subModel.View() + m.helpView()
	}
	return gap + "            Welcome to T-Notes!\n" + gap + m.helpView() 
}
