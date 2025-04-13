package hetzner

import (
	"context"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func CreateVM(ctx context.Context, client *hcloud.Client, name string) (*hcloud.Server, error) {
	serverType := "cx21"              // Small dev server
	image := "ubuntu-22.04"           // Good default
	location := "nbg1"                // Germany DC
	sshKeyName := "your-ssh-key-name" // SSH key pre-loaded in Hetzner panel

	opts := hcloud.ServerCreateOpts{
		Name:       name,
		ServerType: &hcloud.ServerType{Name: serverType},
		Image:      &hcloud.Image{Name: image},
		Location:   &hcloud.Location{Name: location},
		SSHKeys: []*hcloud.SSHKey{
			{Name: sshKeyName},
		},
	}

	result, _, err := client.Server.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return result.Server, nil
}
