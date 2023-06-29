package install

import (
	"os/exec"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

func RunPrecommitInstall() error {
	logger.LogStart("RunPrecommitInstall")
	cmd := exec.Command("pre-commit", "install", "--install-hooks")

	err := cmd.Run()
	if err != nil {
		logger.LogError("Failed to run pre-commit install --install-hooks", err)
		return err
	}
	logger.LogStop("RunPrecommitInstall")
	return nil
}
