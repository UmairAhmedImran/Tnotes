package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Screen int

const gap = "\n\n"

var content = gap + "     Welcome to Tnotes" + gap

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
	viewport viewport.Model
}

//var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("41")).Render

//func helpView() string {
//	return helpStyle("            Hello")
//} not using at the moment

func BaseScreen() (*Model, error) {

	const width = 78

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	const glamourGutter = 2
	glamourRenderWidth := width - vp.Style.GetHorizontalFrameSize() - glamourGutter

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(glamourRenderWidth),
	)

	if err != nil {
		return nil, err
	}
	str, err := renderer.Render(content)
	if err != nil {
		return nil, err
	}
	vp.SetContent(str)
	return &Model{
		viewport: vp,
	}, nil
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
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
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
		return m.viewport.View()
	}
	return m.viewport.View()
}
