package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Devaraja-Anu/ossifix/internal/models"
	"github.com/Devaraja-Anu/ossifix/ui/selector"
	"github.com/Devaraja-Anu/ossifix/ui/spinner"
	textinput "github.com/Devaraja-Anu/ossifix/ui/textInput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initalize new project interactively",
	Run: func(cmd *cobra.Command, args []string) {

		var projectData models.ProjectDetails

		projectInitModel := textinput.InitModel("Enter project init name", "init-name-here")
		projectInitProgram := tea.NewProgram(projectInitModel)

		finalProjectName, err := projectInitProgram.Run()
		if err != nil {
			fmt.Println("Error generating prompt", err)
			os.Exit(1)
		}

		if m, ok := finalProjectName.(textinput.Model); ok {
			projectName := strings.TrimSpace(m.TextInput.Value())
			projectName = strings.ReplaceAll(projectName, " ", "-")

			if projectName == "" {
				fmt.Println("Root name cannot be empty. Aborting")
				os.Exit(1)
			}
			projectData.ProjectName = projectName
		}

		folderNameModel := textinput.InitModel("Enter root folder name", "root-folder-name")
		program := tea.NewProgram(folderNameModel)

		finalFolderNameModel, err := program.Run()
		if err != nil {
			fmt.Println("Error generating prompt", err)
			os.Exit(1)
		}

		if m, ok := finalFolderNameModel.(textinput.Model); ok {
			folderName := m.TextInput.Value()
			if folderName == "" {
				fmt.Println("Project name cannot be empty. Aborting")
				os.Exit(1)
			}

			projectData.RootName = folderName
		}

		optionList := selector.NewSelector("Select router", []string{
			"Gin", "Chi", "Fiber", "Echo", "http",
		})

		selectedOption := tea.NewProgram(optionList)
		finalOptionModal, err := selectedOption.Run()
		if err != nil {
			fmt.Println("Error selecting router")
			os.Exit(1)
		}

		if m, ok := finalOptionModal.(selector.Model); ok {
			selectedChoice := m.Choice
			projectData.Router = selectedChoice
		}

		p := tea.NewProgram(spinner.NewModel(projectData))
		if _, err := p.Run(); err != nil {
			fmt.Println("Error during scaffolding init", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
