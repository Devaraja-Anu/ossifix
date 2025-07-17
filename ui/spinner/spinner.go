package spinner

import (
	"fmt"
	"os/exec"

	"github.com/Devaraja-Anu/ossifix/internal/scaffold"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	steps    []string
	index    int
	spinner  spinner.Model
	done     bool
	rootpath string
}

var stepsList = []string{
	"Initializing project",
	"Creating files",
	"Installing Dependencies",
}

func NewModel(rootpath string) model {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		steps:    stepsList,
		rootpath: rootpath,
		spinner:  s,
	}
}

type completedMessage string

// runstep runs actual logic

func (m model) Init() tea.Cmd {
	return tea.Batch(runSteps(m.index, m.rootpath), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	checkMark := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")

	switch msg := msg.(type) {
	case completedMessage:
		fmt.Println(checkMark.String(), m.steps[m.index])
		m.index++
		if m.index >= len(m.steps) {
			m.done = true
			return m, tea.Quit
		}
		return m, runSteps(m.index, m.rootpath)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return lipgloss.NewStyle().MarginTop(1).Render("✓ Project setup complete!\n")
	}
	return fmt.Sprintf(
		"%s %s...",
		m.spinner.View(),
		lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(m.steps[m.index]),
	)
}

func runSteps(stepIndex int, rootpath string) tea.Cmd {
	return func() tea.Msg {
		var err error
		switch stepIndex {
		case 0:
			cmd := exec.Command("go", "mod", "init", rootpath)
			cmd.Dir = rootpath
			err = cmd.Run()
			if err != nil {
				return completedMessage("✗ Failed to init project")
			}
		case 1:
			err = scaffold.CreateFiles(rootpath)
			if err != nil {
				return completedMessage("✗ Failed to parse templates ")
			}
		case 2:
			mod := exec.Command("go", "mod", "tidy")
			mod.Dir = rootpath
			err = mod.Run()
			if err != nil {
				return completedMessage("✗ Failed to install all dependencies ")
			}
		}

		if err != nil {
			return completedMessage("Encounted an Error!! :( ")
		}
		return completedMessage(fmt.Sprintf("✓ %s completed", stepsList[stepIndex]))
	}
}
