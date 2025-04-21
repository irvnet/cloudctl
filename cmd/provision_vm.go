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

		server, _, err := client.Server.GetByName(ctx, vmName)
		if err != nil {
			fmt.Printf("Err: finding %s by name: %s", vmName, err)
		}

		if server != nil {
			fmt.Printf("ID:%v: / %v / %v already created", server.ID, server.Name, server.Status)
		} else {
			server, err := hetzner.CreateVM(ctx, client, vmName)
			if err != nil {
				fmt.Printf("❌ Failed to create VM: %v\n", err)
				return
			}
			fmt.Printf("✅ Created VM: %s (ID: %d)\n", server.Name, server.ID)

		}
	},
}

func init() {
	provisionCmd.AddCommand(provisionVmCmd)
	provisionVmCmd.Flags().StringVarP(&vmName, "name", "n", "", "Name of the VM to create")
	provisionVmCmd.MarkFlagRequired("name")
}
