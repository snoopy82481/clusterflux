package install

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

func VerifyAndInstallPackages(packages []string) error {
	for _, pkg := range packages {
		if !checkIfPackageIsInstalled(pkg) {
			err := installPackage(pkg)
			if err != nil {
				logger.LogError("Failed to install package "+pkg, err)
			} else {
				logger.LogSuccess("Successfully installed package " + pkg)
			}
		} else {
			logger.LogSuccess("Package " + pkg + " is already installed")
		}
	}
	return nil
}

func checkIfPackageIsInstalled(pkg string) bool {
	cmd := exec.Command("brew", "list", pkg)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func installPackage(pkg string) error {
	cmd := exec.Command("brew", "install", pkg)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func GetPackages() ([]string, error) {
	logger.LogStart("GetPackages")
	data, err := os.ReadFile("packages.txt")
	if err != nil {
		logger.LogError("Failed to read packages.txt", err)
		return nil, fmt.Errorf("failed to read packages.txt: %w", err)
	}
	logger.LogStop("GetPackages")
	return strings.Split(string(data), "\n"), nil
}
