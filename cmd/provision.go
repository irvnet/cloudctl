package cmd

import (
	"github.com/spf13/cobra"
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "provision - provision cloud resources",
}

func init() {
	rootCmd.AddCommand(provisionCmd)
}
