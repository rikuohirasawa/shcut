package commands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	internals "github.com/rikuohirasawa/shcut/internals"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Add(configFilePath string) *cobra.Command {
	return &cobra.Command{
		Use:   "add [name] [command]",
		Short: "Add a shortcut (no args → prompt mode)",
		Args:  cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var name, command string
			reader := bufio.NewReader(os.Stdin)

			if len(args) > 2 {
				return errors.New("usage: shcut add [name] [command]")
			}

			switch len(args) {
			case 2:
				name, command = args[0], args[1]
			case 0:
				fmt.Print("Alias name: ")
				line, _ := reader.ReadString('\n')
				name = strings.TrimSpace(line)
				if name == "" {
					return errors.New("alias name cannot be empty")
				}

				fmt.Println("Enter command (Ctrl+D to finish):")
				bytes, err := io.ReadAll(reader)
				if err != nil && err != io.EOF {
					return fmt.Errorf("failed reading command: %w", err)
				}
				command = strings.TrimSpace(string(bytes))
			default:
				return errors.New("usage: shcut add [name] [command]")
			}

			config, err := internals.LoadConfig(configFilePath)
			if err != nil {
				return err
			}

			if _, exists := config[name]; exists {
				fmt.Printf("Alias '%s' exists – overwrite? (y/N): ", name)
				ans, _ := reader.ReadString('\n')
				ans = strings.TrimSpace(strings.ToLower(ans))
				if ans != "y" && ans != "yes" {
					log.Info("aborted")
					return nil
				}
			}

			config[name] = command
			if err := internals.SaveConfig(config, configFilePath); err != nil {
				return err
			}
			log.Infof("saved: %s -> %s", name, command)
			return nil
		},
	}
}
