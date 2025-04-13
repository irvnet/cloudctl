package hetzner

import (
	"github.com/hetznercloud/hcloud-go/hcloud"
	"os"
)

func NewClient() *hcloud.Client {
	token := os.Getenv("HCLOUD_TOKEN")
	if token == "" {
		panic("HCLOUD_TOKEN environment variable not set")
	}
	return hcloud.NewClient(hcloud.WithToken(token))
}
