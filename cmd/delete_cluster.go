package cmd

import (
	"cloudctl/internal/infra"
	"fmt"
	"github.com/spf13/cobra"
)

var deleteClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Delete the cluster (ctrl + workers)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.DeleteComponent("worker", 3); err != nil {
			return fmt.Errorf("failed to delete workers: %w", err)
		}

		if err := infra.DeleteComponent("ctrl", 1); err != nil {
			return fmt.Errorf("failed deleting ctrl %w", err)
		}

		fmt.Println("[cloudctl] Done deleting cluster (ctrl + workers)...")
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterCmd)
}
