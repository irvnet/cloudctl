package cmd

import (
	"cloudctl/internal/infra"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [component]",
	Short: "Create a component (devbox, ctrl, worker)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		component := args[0]
		count, _ := cmd.Flags().GetInt("count")
		return infra.CreateComponent(component, count)

	},
}

func init() {
	createCmd.Flags().Int("count", 2, "Number of workers to create")
	rootCmd.AddCommand(createCmd)
}
