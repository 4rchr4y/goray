package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of the application",
	Long: `The version command is used to display the current version of the application.
It provides the detailed version information including the major, minor, and patch numbers.`,
	Run: runVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("version 0.0.1")
}
