package commands

import (
	"fmt"

	"os"
	"os/exec"
	"strings"

	internals "github.com/rikuohirasawa/shcut/internals"

	"github.com/spf13/cobra"
)

func Run(configFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:   "run [name] [--] [args...]",
		Short: "Execute a defined shortcut",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			config, err := internals.LoadConfig(configFilePath)
			if err != nil {
				return err
			}
			command, exists := config[name]
			if !exists {
				return fmt.Errorf("shortcut '%s' not found", name)
			}
			if len(args) > 1 {
				command = command + " " + strings.Join(args[1:], " ")
			}
			// TODO: make this dynamic
			c := exec.Command("/bin/sh", "-c", command)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		},
	}
}
