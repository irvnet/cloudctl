package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a component (devbox, ctrl, worker)",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
