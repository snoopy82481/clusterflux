package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/snoopy82481/clusterflux/internal/install"
	"github.com/snoopy82481/clusterflux/internal/logger"
	"github.com/thlib/go-timezone-local/tzlocal"
)

func CreateConfig() (Config, error) {
	logger.LogStart("CreateConfig")
	agePublicKey, err := install.GetAgePublicKey()
	if err != nil {
		logger.LogError("Unable to get AGE public key", err)
		return Config{}, err
	}

	cloudflared, err := getCloudflareCreds()
	if err != nil {
		logger.LogError("Unable to get clouldflard creds", err)
		return Config{}, err
	}

	tzname, err := tzlocal.RuntimeTZ()
	if err != nil {
		logger.LogError("Unable to get local machine timezone", err)
		return Config{}, err
	}

	logger.LogStop("CreateConfig")

	return Config{
		Email:        "",
		Timezone:     tzname,
		AgePublicKey: agePublicKey,
		Apps: Apps{
			WeaveGitOps: WeaveGitOps{
				AdminPassword: "generate",
			},
			Grafana: Grafana{
				AdminPassword: "generate",
			},
		},
		Network: Network{
			ClusterCidr:    "10.42.0.0/16",
			ServiceCidr:    "10.43.0.0/16",
			K8sGatewayAddr: "",
			IngressAddr:    "",
			KubeVIPAddr:    "",
		},
		GitHub: GitHubConfig{
			Public: true,
			URL:    "",
			FluxWebhook: FluxWebhook{
				Secret: "generate",
			},
		},
		Cloudflare: CloudflareConfig{
			Domain:   "",
			APIToken: "",
			Tunnel: CloudflareTunnel{
				AccountTag:   cloudflared.AccountTag,
				TunnelSecret: cloudflared.TunnelSecret,
				TunnelID:     cloudflared.TunnelID,
			},
		},
		Ansible: AnsibleConfig{
			Enabled:                   true,
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.LogError("Failed to retrieve home directory", err)
		return nil, err
	}

	cloudflaredDir := os.ExpandEnv(homeDir + "/.cloudflared")
	var cloudflareFile string

	err = filepath.Walk(cloudflaredDir, func(path string, info os.FileInfo, err error) error {
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
