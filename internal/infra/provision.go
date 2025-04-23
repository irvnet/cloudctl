package infra

import (
	"context"
	"fmt"
	"os"
	"strings"

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

func instanceExists(ctx context.Context, client *hcloud.Client, name string) (bool, error) {
	server, _, err := client.Server.GetByName(ctx, name)
	if err != nil {
		return false, err
	}

	return server != nil, nil

}

func resolveComponentConfig(name string) (ComponentConfig, error) {
	normalized := strings.ToLower(name)
	if normalized == "workers" {
		normalized = "worker"
	}

	cfg, ok := ComponentDefaults[normalized]
	if !ok {
		return ComponentConfig{}, fmt.Errorf("unknown component: %s", name)
	}
	return cfg, nil
}
