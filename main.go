package main

import (
	"fmt"
	"os"
	"path/filepath"

	commands "github.com/rikuohirasawa/shcut/commands"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFilePath = "$HOME/.shcut/config.json"

var rootCmd = &cobra.Command{
	Use:   "shcut",
	Short: "Manage your shell shortcuts",
	Long:  `shcut is a CLI tool to add, list, remove, and run shell shortcuts easily.`,
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to find home directory: %v\n", err)
		os.Exit(1)
	}
	configFilePath = filepath.Join(home, ".shcut", "config.json")

	dir := filepath.Dir(configFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create config directory: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		if err := os.WriteFile(configFilePath, []byte("{}"), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create alias file: %v\n", err)
			os.Exit(1)
		}
	}
}

func main() {
	rootCmd.AddCommand(commands.Add(configFilePath), commands.Remove(configFilePath), commands.Run(configFilePath), commands.Browse(configFilePath))
	rootCmd.Execute()
}
