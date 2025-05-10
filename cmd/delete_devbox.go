package cmd

import (
	"cloudctl/internal/infra"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteDevboxCmd = &cobra.Command{
	Use:   "devbox",
	Short: "Delete the devbox vm",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.DeleteComponent("devbox", 1); err != nil {
			return fmt.Errorf("failed to delete devbox: %w", err)
		}

		fmt.Println("[cloudctl] Devbox deleted successfully")
		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteDevboxCmd)
}
