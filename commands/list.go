package commands

import (
	internals "github.com/rikuohirasawa/shcut/internals"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func List(configFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all defined shortcuts",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := internals.LoadConfig(configFilePath)
			if err != nil {
				return err
			}
			if len(config) == 0 {
				log.Info("No shortcuts defined.")
				return nil
			}
			log.Info("Shortcuts:")
			for name, command := range config {
				log.Infof("  %s -> %s", name, command)
			}
			return nil
		},
	}
}
