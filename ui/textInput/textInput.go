package textinput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Model struct {
	TextInput textinput.Model
	err       error
}

func InitModel() Model {
	ti := textinput.New()
	ti.Placeholder = "your-project-name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		TextInput: ti,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"What is the name of the root folder?\n\n%s\n\n%s",
		m.TextInput.View(),
		"(esc to quit)",
	)
}
