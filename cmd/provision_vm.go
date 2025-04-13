package cmd

import (
	"cloudctl/pkg/hetzner"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var vmName string

var provisionVmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Provision a single VM",
	Run: func(cmd *cobra.Command, args []string) {
		client := hetzner.NewClient()
		ctx := context.Background()

		server, err := hetzner.CreateVM(ctx, client, vmName)
		if err != nil {
			fmt.Printf("❌ Failed to create VM: %v\n", err)
			return
		}

		fmt.Printf("✅ Created VM: %s (ID: %d)\n", server.Name, server.ID)
	},
}

func init() {
	provisionCmd.AddCommand(provisionVmCmd)
	provisionVmCmd.Flags().StringVarP(&vmName, "name", "n", "", "Name of the VM to create")
	provisionVmCmd.MarkFlagRequired("name")
}
