package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Screen int

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
  return helpStyle("Hello")
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
		case "q":
			m.quitting = true
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
		return m.subModel.View() + m.helpView()
	}
	return "Welcome to T-Notes!\n[q] to quit. " + m.helpView()
}

//import (
//	"fmt"
//
//	tea "github.com/charmbracelet/bubbletea"
//	"github.com/charmbracelet/lipgloss"
//)
//
//var (
//	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235"))
//	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
//)
//
//type Model struct {
//	altscreen  bool
//	quitting   bool
//	suspending bool
//}
//
//func (m Model) Init() tea.Cmd {
//	return nil
//}
//
//func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	switch msg := msg.(type) {
//	case tea.ResumeMsg:
//		m.suspending = false
//		return m, nil
//	case tea.KeyMsg:
//		switch msg.String() {
//		case "q", "ctrl+c", "esc":
//			m.quitting = true
//			return m, tea.Quit
//		case "ctrl+z":
//			m.suspending = true
//			return m, tea.Suspend
//		case " ":
//			var cmd tea.Cmd
//			if m.altscreen {
//				cmd = tea.ExitAltScreen
//			} else {
//				cmd = tea.EnterAltScreen
//			}
//			m.altscreen = !m.altscreen
//			return m, cmd
//		}
//	}
//	return m, nil
//}
//
//func (m Model) View() string {
//	if m.suspending {
//		return ""
//	}
//
//	if m.quitting {
//		return "Bye!\n"
//	}
//
//	const (
//		altscreenMode = " altscreen mode "
//		inlineMode    = " inline mode "
//	)
//
//	var mode string
//	if m.altscreen {
//		mode = altscreenMode
//	} else {
//		mode = inlineMode
//	}
//
//	return fmt.Sprintf("\n\n  You're in %s\n\n\n", keywordStyle.Render(mode)) +
//		helpStyle.Render("  space: switch modes • ctrl-z: suspend • q: exit\n")
//}
//
