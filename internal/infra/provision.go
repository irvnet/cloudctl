package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func CreateComponent(name string, count int) error {
	cfg, err := resolveComponentConfig(name)
	if err != nil {
		return err
	}

	client := hcloud.NewClient(hcloud.WithToken(os.Getenv("HCLOUD_TOKEN")))
	ctx := context.Background()
	instances := generateInstanceNames(name, count)

	for _, instanceName := range instances {
		exists, err := instanceExists(ctx, client, instanceName)
		if err != nil {
			return fmt.Errorf("failed to cehck instance existance %w", err)
		}
		if exists {
			fmt.Printf("⚠️  %s already exists, skipping...\n", instanceName)
			continue
		}

		if err := createServer(); err != nil {
			return fmt.Errorf("failed to create %s: %w", instanceName, err)
		}
	}

	return nil

}
