/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/snoopy82481/clusterflux/internal/config"
	"github.com/snoopy82481/clusterflux/internal/logger"
	"github.com/snoopy82481/clusterflux/internal/validation"
	"github.com/spf13/cobra"
)

var (
	k      = koanf.New(".")
	parser = yaml.Parser()
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates files for use in kubernetes cluster",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("logfile.log")
		if os.IsNotExist(err) {
			var file, err = os.Create("logfile.log")
			if err != nil {
				log.Panicf("Unable to Create logfile. %v", err)
				return
			}

			defer file.Close()
		}

		logfile, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return
		}
		defer logfile.Close()

		logger.LogInit(false, logfile)

		defer func() {
			_ = logger.GetLogger().Sync()
		}()

		logger.LogStart("generate")

		if err := k.Load(file.Provider("config.yaml"), parser); err != nil {
			logger.LogError("Failed to load config file", err)
			return
		}

		config := &config.Config{}
		if err := k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
			logger.LogError("Failed to unmarshal config:", err)
			return
		}

		if err := validation.ValidateConfig(config); err != nil {
			logger.LogError("", err)
			return
		}

		if err := validation.ValidateAnsibleHosts(config); err != nil {
			logger.LogError("Unable to validate Ansible Hosts", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
