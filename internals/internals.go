package internals

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

func LoadConfig(aliasFile string) (map[string]string, error) {
	data, err := os.ReadFile(aliasFile)
	if err != nil {
		return nil, err
	}
	config := make(map[string]string)

	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Error("Unable to unmarshal config:", err)
		return nil, err
	}

	return config, nil
}

func SaveConfig(config map[string]string, aliasFile string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(aliasFile, data, 0644)
}
