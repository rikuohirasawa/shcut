package commands

import (
	"fmt"

	internals "github.com/rikuohirasawa/shcut/internals"
	"github.com/spf13/cobra"
)

func Remove(configFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove", "delete"},
		Short:   "Remove an existing shortcut",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			aliases, err := internals.LoadConfig(configFilePath)
			if err != nil {
				return err
			}
			if _, exists := aliases[name]; !exists {
				fmt.Printf("shortcut '%s' does not exist", name)
				return nil
			}
			delete(aliases, name)
			if err := internals.SaveConfig(aliases, configFilePath); err != nil {
				return err
			}
			return nil
		},
	}
}
