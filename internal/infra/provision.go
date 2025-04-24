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

	token := os.Getenv("HCLOUD_TOKEN")
	client := hcloud.NewClient(hcloud.WithToken(token))
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

		if err := createServer(ctx, client, cfg, name); err != nil {
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

func createServer(ctx context.Context, client *hcloud.Client, cfg ComponentConfig, name string) error {

	key, _, err := client.SSHKey.GetByName(ctx, cfg.SSHKey)
	if err != nil {
		return fmt.Errorf("unable to find SSH key %q: %w", cfg.SSHKey, err)
	}
	if key == nil {
		return fmt.Errorf("SSH key %q not found", cfg.SSHKey)
	}

	opts := hcloud.ServerCreateOpts{
		Name:             name,
		ServerType:       &hcloud.ServerType{Name: cfg.ServerType},
		Image:            &hcloud.Image{Name: cfg.Image},
		Location:         &hcloud.Location{Name: cfg.Location},
		SSHKeys:          []*hcloud.SSHKey{key},
		Labels:           cfg.Labels,
		StartAfterCreate: hcloud.Ptr(true),
	}

	resp, _, err := client.Server.Create(ctx, opts)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Created %s (ID: %d, IP: %s)\n", name, resp.Server.Name, resp.Server.PublicNet.IPv4.IP)
	return nil

}
