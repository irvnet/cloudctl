package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete infrastructure components (devbox, ctrl, workers)",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	fmt.Println("running delete init...")
	rootCmd.AddCommand(deleteCmd)
}
