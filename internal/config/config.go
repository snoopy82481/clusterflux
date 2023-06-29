package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/snoopy82481/clusterflux/internal/install"
	"github.com/snoopy82481/clusterflux/internal/logger"
)

func CreateConfig() (Config, error) {
	logger.LogStart("CreateConfig")
	agePublicKey, err := install.GetAgePublicKey()
	if err != nil {
		logger.LogError("Unable to get AGE public key", err)
		return Config{}, err
	}

	cloudflaredOutput, err := getCloudflareCreds()
	if err != nil {
		logger.LogError("Unable to get clouldflard creds", err)
	}

	logger.LogStop("CreateConfig")
	return Config{
		Core: CoreConfig{
			MetalLBRange:             "",
			AgePublicKey:             agePublicKey,
			Timezone:                 time.Now().Location().String(),
			WeaveGitOpsAdminPassword: "generate",
		},
		GitHub: GitHubConfig{
			URL:               "",
			FluxWebhookSecret: "generate",
		},
		Cloudflare: CloudflareConfig{
			Domain: "",
			Email:  "",
			APIKey: "",
			Tunnel: CloudflareTunnel{
				AccountTag:   cloudflaredOutput.AccountTag,
				TunnelSecret: cloudflaredOutput.TunnelSecret,
				TunnelID:     cloudflaredOutput.TunnelID,
			},
		},
		Ansible: AnsibleConfig{
			ControlNodeHostnamePrefix: "control",
			NodeHostnamePrefix:        "node",
			Hosts: []AnsibleHost{
				{
					IPAddress:    "",
					SSHUsername:  "",
					SudoPassword: "",
					ControlNode:  true,
					Hostname:     "",
				},
			},
		},
	}, nil
}

func getCloudflareCreds() (*CloudflareTunnel, error) {
	cloudflaredDir := os.ExpandEnv("~/.cloudflared")
	var cloudflareFile string

	err := filepath.Walk(cloudflaredDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.LogError("Failed to access path "+path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			cloudflareFile = path
		}
		return nil
	})

	if err != nil {
		logger.LogError("Failed to walk cloudflared directory", err)
		return nil, fmt.Errorf("failed to walk cloudflared directory: %w", err)
	}
	if cloudflareFile == "" {
		logger.LogWarn("No .json file found in cloudflared directory")
		return nil, fmt.Errorf("no .json file found in cloudflared directory")
	}

	data, err := os.ReadFile(cloudflareFile)
	if err != nil {
		logger.LogError("Failed to read Cloudflare JSON file", err)
		return nil, fmt.Errorf("failed to read Cloudflare JSON file: %w", err)
	}

	var config CloudflareTunnel
	err = json.Unmarshal(data, &config)
	if err != nil {
		logger.LogError("Failed to unmarshal Cloudflare JSON file", err)
		return nil, fmt.Errorf("failed to unmarshal Cloudflare JSON file: %w", err)
	}

	logger.LogInfo("Successfully retrieved Cloudflare configuration", "GetCloudflareConfig")
	return &config, nil
}
