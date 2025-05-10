package cmd

import (
	"cloudctl/internal/infra"
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCtrlCmd = &cobra.Command{
	Use:   "ctrl",
	Short: "Delete the k8s control plane vm",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.DeleteComponent("ctrl", 1); err != nil {
			return fmt.Errorf("failed to delete ctrl %w", err)
		}

		fmt.Println("[cloudctl] Control plane deleted successfully")
		return nil

	},
}

func init() {
	deleteCmd.AddCommand(deleteCtrlCmd)
}
