package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:"ossifix",
	Short:"A modular CLi scaffolding tool with cobra and bubbletea",
}

func Execute() error {
	return rootCmd.Execute()
}

