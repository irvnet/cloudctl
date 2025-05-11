package cmd

import (
	"cloudctl/internal/infra"
	"fmt"

	"github.com/spf13/cobra"
)

var createDevboxCmd = &cobra.Command{
	Use:   "devbox",
	Short: "Create a temp workstation",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := infra.CreateComponent("devbox", 1); err != nil {
			return fmt.Errorf("devbox creation failed: %w", err)
		}

		//home	if err := infra.RunBootstrapScript("devbox"); err != nil {
		//		return fmt.Errorf("bootstrap failed: %w", err)
		//	}

		fmt.Println("[cloudctl] Devbox created successfully")
		return nil
	},
}

func init() {
	createCmd.AddCommand(createDevboxCmd)
}
