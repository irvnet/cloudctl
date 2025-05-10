package cmd

import (
	"cloudctl/internal/infra"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteWorkersCmd = &cobra.Command{
	Use:   "workers",
	Short: "Delete all worker nodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.DeleteComponent("worker", 3); err != nil {
			return fmt.Errorf("failed to delete workers: %w", err)
		}

		fmt.Println("[cloudctl] Done deleting worker nodes...")
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteWorkersCmd)
}
