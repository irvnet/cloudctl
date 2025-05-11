package cmd

import (
	"cloudctl/internal/infra"
	"fmt"

	"github.com/spf13/cobra"
)

var createCtrlCmd = &cobra.Command{
	Use:   "ctrl",
	Short: "Create k8s control plane node",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.CreateComponent("ctrl", 1); err != nil {
			return fmt.Errorf("ctrl creation failed: %w", err)
		}

		//	if err := infra.RunBootstrapScript("ctrl"); err != nil {
		//		return fmt.Errorf("bootstrap failed %w", err)
		//	}

		fmt.Println("[cloudctl] Control plane created.")
		return nil
	},
}

func init() {
	createCmd.AddCommand(createCtrlCmd)
}
