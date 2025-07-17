package cmd

import (
	"fmt"
	"os"

	"github.com/Devaraja-Anu/ossifix/ui/spinner"
	textinput "github.com/Devaraja-Anu/ossifix/ui/textInput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize new project interactively",
	Run: func(cmd *cobra.Command, args []string) {
		model := textinput.InitModel()
		program := tea.NewProgram(model)

		finalModel, err := program.Run()
		if err != nil {
			fmt.Println("Error generating prompt", err)
			os.Exit(1)
		}

		if m, ok := finalModel.(textinput.Model); ok {
			projectName := m.TextInput.Value()
			if projectName == "" {
				fmt.Println("Project name cannot be empty. Aborting")
				os.Exit(1)
			}

			p := tea.NewProgram(spinner.NewModel(projectName))
			if _, err := p.Run(); err != nil {
				fmt.Println("Error during scaffolding init", err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
