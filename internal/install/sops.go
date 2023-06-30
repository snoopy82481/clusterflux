package install

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

func GetAgePublicKey() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.LogError("Failed to retrieve home directory", err)
		return "", err
	}

	keysFile := filepath.Join(homeDir, ".config/sops/age/keys.txt")
	data, err := os.ReadFile(keysFile)
	if err != nil {
		logger.LogError("Failed to read keys file", err)
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		logger.LogWarn("Unexpected keys file format")
		return "", err
	}

	parts := strings.Split(lines[1], ":")
	if len(parts) < 2 {
		logger.LogWarn("Unexpected keys file format")
		return "", err
	}

	publicKey := strings.TrimSpace(parts[1])

	logger.LogInfo("Successfully retrieved public key", "GetAgePublicKey")
	return publicKey, nil
}

func SetupSopsAge() error {
	logger.LogStart("SetupSopsAge")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.LogError("Failed to retrieve home directory", err)
		return err
	}

	sopsDir := filepath.Join(homeDir, ".config/sops/age")
	keysFile := filepath.Join(sopsDir, "/keys.txt")

	if _, err := os.Stat(sopsDir); os.IsNotExist(err) {
		logger.LogInfo("Creating sops directory...", "SetupSopsAge")

		err := os.MkdirAll(sopsDir, 0755)
		if err != nil {
			logger.LogError("Failed to create directory", err)
			return err
		}
	}

	if _, err := os.Stat(keysFile); os.IsNotExist(err) {
		logger.LogInfo("Running age-keygen command...", "SetupSopsAge")

		cmd := exec.Command("age-keygen", "-o", keysFile)
		err := cmd.Run()
		if err != nil {
			logger.LogError("Failed to run age-keygen", err)
			return err
		}
	}

	logger.LogStop("SetupSopsAge")
	return nil
}
