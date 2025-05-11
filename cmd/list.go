package cmd

import (
	"cloudctl/internal/infra"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List provisioned components",
	RunE: func(cmd *cobra.Command, args []string) error {
		return infra.ListComponents()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
