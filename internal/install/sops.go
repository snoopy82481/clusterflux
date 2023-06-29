package install

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/snoopy82481/clusterflux/internal/logger"
)

// GetAgePublicKey function reads the age public key from keys.txt and returns it
func GetAgePublicKey() (string, error) {
	keysFile := os.ExpandEnv("~/.config/sops/age/keys.txt")
	data, err := os.ReadFile(keysFile)
	if err != nil {
		logger.LogError("Failed to read keys file", err)
		return "", fmt.Errorf("failed to read keys file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		logger.LogWarn("Unexpected keys file format")
		return "", fmt.Errorf("unexpected keys file format")
	}

	parts := strings.Split(lines[1], ":")
	if len(parts) < 2 {
		logger.LogWarn("Unexpected keys file format")
		return "", fmt.Errorf("unexpected keys file format")
	}

	publicKey := strings.TrimSpace(parts[1])

	logger.LogInfo("Successfully retrieved public key", "GetAgePublicKey")
	return publicKey, nil
}

func SetupSopsAge() error {
	logger.LogStart("SetupSopsAge")

	sopsDir := os.ExpandEnv("~/.config/sops/age")
	keysFile := os.ExpandEnv(sopsDir + "/keys.txt")

	if _, err := os.Stat(sopsDir); os.IsNotExist(err) {
		logger.LogInfo("Creating sops directory...", "SetupSopsAge")

		err := os.MkdirAll(sopsDir, 0755)
		if err != nil {
			logger.LogError("Failed to create directory", err)
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	if _, err := os.Stat(keysFile); os.IsNotExist(err) {
		logger.LogInfo("Running age-keygen command...", "SetupSopsAge")

		cmd := exec.Command("age-keygen", "-o", keysFile)
		err := cmd.Run()
		if err != nil {
			logger.LogError("Failed to run age-keygen", err)
			return fmt.Errorf("failed to run age-keygen: %w", err)
		}
	}

	_, err := GetAgePublicKey()
	if err != nil {
		return fmt.Errorf("failed to get public key: %w", err)
	}

	logger.LogStop("SetupSopsAge")
	return nil
}
