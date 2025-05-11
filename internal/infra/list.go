package infra

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func ListComponents() error {

	const (
		colorGreen = "\033[32m"
		colorReset = "\033[0m"
	)

	token := os.Getenv("HCLOUD_TOKEN")
	client := hcloud.NewClient(hcloud.WithToken(token))
	ctx := context.Background()

	servers, err := client.Server.All(ctx)
	if err != nil {
		return fmt.Errorf("failed to list servers: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 8, ' ', 0)
	fmt.Fprintln(w, "SERVER\t ROLE\t IP\t LOCATION")

	for _, server := range servers {
		role := server.Labels["role"]
		ip := server.PublicNet.IPv4.IP
		location := server.Datacenter.Location.Name

		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", server.Name, role, ip, location)

	}

	w.Flush()
	return nil
}
