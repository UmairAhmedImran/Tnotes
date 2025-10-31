package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	input string
	done  bool
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.done = true
			// You might send a message back to root model to save
			return m, tea.Quit
		case "esc":
			return m, tea.Quit
		default:
			if len(msg.String()) == 1 {
				m.input += msg.String()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.done {
		return "Note saved! Press any key to return.\n"
	}
	return "Add a note:\n[" + m.input + "]\n(Press Enter to save, Esc to cancel)\n"
}

func (m Model) Value() string {
	return m.input
}