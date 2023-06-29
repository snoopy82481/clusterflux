package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/snoopy82481/clusterflux/internal/config"
	"github.com/snoopy82481/clusterflux/internal/install"
	"github.com/snoopy82481/clusterflux/internal/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes dev envirionment",
	Long: `Initializes the development envirionment.

	It will then do the following steps:
	1. Verify/Install brew from brew.sh
	2. Verify/Install software dependencies listed in the init task of https://raw.githubusercontent.com/onedr0p/flux-cluster-template/main/Taskfile.yml preferably listed in an external file for easy updating
	3. Run "pre-commit install --install-hooks"
	4. Run "mkdir-p ~/.config/sops/age" then "age-keygen -o ~/.config/sops/keys.txt"
	5. Run "cloudflared tunnel login" then "cloudflared tunnel create k8s"
	6. Create config.yaml where it does the following:
		- Set fluxGitHubWebhookSecret = generate
		- Set weaveGitOpsAdminPassword = generate
		- Set controlNodeHostnamePrefix = control
		- Set nodeHostnamePrefix = node
		- Set timezone to local machine Timezone
		- Set agePublicKey from ~/.config/sops/keys.txt where it gets the second line and gets everything past the : and trims the spaces
		- Set accountTag,tunnelSecret,tunnelID from json file in ~/.cloudflared folder
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("logfile.log")
		if os.IsNotExist(err) {
			var file, err = os.Create("logfile.log")
			if err != nil {
				log.Panicf("Unable to Create logfile. %v", err)
				return
			}
			log.Printf("Thumbs up!")

			time.Sleep(3 * time.Second)

			defer file.Close()
		}

		// If the file exists but there's another error, handle it
		if err != nil {
			log.Panicf("Failed to access logfile.log. %v", err)
			return
		}

		logfile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return
		}
		defer logfile.Close()

		logger.LogInit(false, logfile)

		defer func() {
			_ = logger.GetLogger().Sync()
		}()

		logger.LogStart("init")

		err = install.VerifyInstallBrew()
		if err != nil {
			logger.LogError("Failed to Install Brew", err)
			os.Exit(1)
		}

		packages, err := install.GetPackages()
		if err != nil {
			logger.LogError("Failed to get packages", err)
			return
		}

		err = install.VerifyAndInstallPackages(packages)
		if err != nil {
			logger.LogError("Failed to install packages", err)
			return
		}

		logger.LogSuccess("All packages installed successfully")

		err = install.RunPrecommitInstall()
		if err != nil {
			logger.LogError("Failed to install Pre-commit Hooks", err)
			return
		}

		logger.LogSuccess("Pre-Commit installed successfully")

		err = install.SetupSopsAge()
		if err != nil {
			logger.LogError("Failed to setup Age", err)
			return
		}

		publicKey, err := install.GetAgePublicKey()
		if err != nil {
			logger.LogError("Failed to get Age public key", err)
			return
		}

		logger.LogInfo("Public Key: "+publicKey, "init")

		err = install.CloudflaredLoginAndCreateTunnel()
		if err != nil {
			logger.LogError("Failed to setup Cloudflared tunnel", err)
			return
		}

		logger.LogSuccess("Cloudflared tunnel setup.")

		config, err := config.CreateConfig()
		if err != nil {
			logger.LogError("Failed to create configuration", err)
			return
		}

		err = writeConfigToFile(config, "config.yaml")
		if err != nil {
			logger.LogError("Failed to write configuration to file", err)
		}

		logger.LogSuccess("Configuration initialized successfully")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func writeConfigToFile(configInput config.Config, filename string) error {
	data, err := yaml.Marshal(configInput)
	if err != nil {
		logger.LogError("Failed to marsha config to YAML", err)
		return fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		logger.LogError("Failed to create config file", err)
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		logger.LogError("Failed to write config to file", err)
		return fmt.Errorf("failed to write config to file: %w", err)
	}

	logger.LogInfo("Successfully wrote config to file", "WriteConfigToFile")
	return nil
}
