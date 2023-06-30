package install

import (
	"os/exec"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

func VerifyInstallBrew() error {
	logger.LogStart("VerifyInstallBrew")

	_, err := exec.LookPath("brew")

	if err != nil {
		logger.LogWarn("Brew not installed, installing now.")

		err := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)").Run()

		if err != nil {
			logger.LogError("failed to install brew: %v", err)
			return err
		}
	}

	logger.LogInfo("Brew is installed", "VerifyInstallBrew")
	logger.LogStop("VerifyInstallBrew")
	return nil
}
