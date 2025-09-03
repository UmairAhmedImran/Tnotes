package addmodel

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
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.done {
		return "Note saved! Press any key to return.\n"
	}
	return "Add a note:\n(Type your note and press Enter)\n"
}
