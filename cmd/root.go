package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudctl",
	Short: "cloudctl - your custom cloud native management CLI",
	Long:  "cloudctl -  helps provision and manage VMs and clusters for your cloud native projects.",
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

