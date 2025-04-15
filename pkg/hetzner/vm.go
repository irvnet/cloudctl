package hetzner

import (
	"context"
	"fmt"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func CreateVM(ctx context.Context, client *hcloud.Client, name string) (*hcloud.Server, error) {
	serverType := "cax11"       // Small dev server
	image := "ubuntu-22.04"     // Good default
	location := "nbg1"          // Germany DC
	sshKeyName := "hetzner-ssh" // SSH key pre-loaded in Hetzner panel

	sshKey, _, err := client.SSHKey.GetByName(ctx, sshKeyName)
	if err != nil || sshKey == nil {
		return nil, fmt.Errorf("SSH key '%s' not found in your account...", sshKeyName)
	}

	opts := hcloud.ServerCreateOpts{
		Name:       name,
		ServerType: &hcloud.ServerType{Name: serverType},
		Image:      &hcloud.Image{Name: image},
		Location:   &hcloud.Location{Name: location},
		SSHKeys:    []*hcloud.SSHKey{sshKey},
	}

	result, _, err := client.Server.Create(ctx, opts)
	if err != nil {
		return nil, err
	}
	return result.Server, nil
}
