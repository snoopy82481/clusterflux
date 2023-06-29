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
		err := checkAndInstallPackage(pkg)
		if err != nil {
			logger.LogError("Failed to install package", err)
			return err
		}
		logger.LogSuccess("Successfully installed package " + pkg)
	}
	return nil
}

func checkAndInstallPackage(pkg string) error {
	if checkIfPackageIsInstalled(pkg) {
		logger.LogInfo(pkg+" is already installed", "checkAndInstallPackage")
	} else {
		err := installPackage(pkg)
		if err != nil {
			logger.LogError("Failed to install package "+pkg, err)
			return err
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
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.LogError("Failed to install package "+pkg+": "+string(output), err)
		return fmt.Errorf("Failed to install package %w", err)
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
