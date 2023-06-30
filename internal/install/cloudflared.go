package install

import (
	"os/exec"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

func CloudflaredLoginAndCreateTunnel() error {
	logger.LogStart("CloudflaredLoginAndCreateTunnel")

	cmd := exec.Command("cloudflared", "tunnel", "login")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.LogError("Failed to execute 'cloudflared tunnel login'", err)
		return err
	}
	logger.LogInfo(string(output), "CloudflaredLoginAndCreateTunnel")

	cmd = exec.Command("cloudflared", "tunnel", "create", "k8s")
	output, err = cmd.CombinedOutput()
	if err != nil {
		logger.LogError("Failed to execute 'cloudflared tunnel create k8s'", err)
		return err
	}
	logger.LogInfo(string(output), "CloudflaredLoginAndCreateTunnel")

	logger.LogStop("CloudflaredLoginAndCreateTunnel")
	return nil
}
