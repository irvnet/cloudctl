package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func DeleteComponent(name string, count int) error {
	token := os.Getenv("HCLOUD_TOKEN")
	client := hcloud.NewClient(hcloud.WithToken(token))
	ctx := context.Background()
	instances := generateInstanceNames(name, count)

	for _, instanceName := range instances {
		server, _, err := client.Server.GetByName(ctx, instanceName)
		if err != nil {
			return fmt.Errorf("failed to lookup %s: %w", instanceName, err)
		}

		if server == nil {
			fmt.Printf("⚠️  %s not found, skipping...\n", instanceName)
			continue
		}

		action, resp, err = client.Server.DeleteWithResult(ctx, server)
		if err != nil {
			return fmt.Errorf("failed to delete %s: %w", instanceName, err)
		}

		fmt.Printf("🗑️  Deleted %s (ID: %d)\n", instanceName, server.ID)
	}

	return nil
}
