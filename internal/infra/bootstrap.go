package infra

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

// validate bootstrap exists before attempting to run it
func scriptExists(name string) (string, error) {
	script := fmt.Sprintf("bootstrap.%s.sh", name)
	if _, err := os.Stat(script); os.IsNotExist(err) {
		return "", fmt.Errorf("script not found: %s", script)
	}

	return script, nil
}

// public ip lookup to enable ssh for bootstrapping
func lookupPublicIP(name string) (string, error) {
	token := os.Getenv("HCLOUD_TOKEN")
	client := hcloud.NewClient(hcloud.WithToken(token))
	ctx := context.Background()

	server, _, err := client.Server.GetByName(ctx, name)
	if err != nil || server == nil {
		return "", fmt.Errorf("server %q not found: %w", name, err)
	}

	ip := server.PublicNet.IPv4.IP
	if ip == nil {
		return "", fmt.Errorf("server %q has no public IPv4 address", name)
	}

	return ip.String(), nil
}

func runScriptOverSSH(scriptPath, ip string) error {
	fmt.Printf("[bootstrap] Running %s on %s\n", scriptPath, ip)
	cmd := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", fmt.Sprintf("root@%s", ip), "bash -s")
	scriptFile, err := os.Open(scriptPath)
	if err != nil {
		return fmt.Errorf("failed to open script: %w", err)
	}

	defer scriptFile.Close()
	cmd.Stdin = scriptFile
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
