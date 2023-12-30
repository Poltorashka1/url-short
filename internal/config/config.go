package config

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Config struct {
	Server         `yaml:"server"`
	DatabaseConfig `yaml:"databaseConfig"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Type   string            `yaml:"type"`
	Config map[string]string `yaml:"config"`
}

// NewConfig creates a new Config object by reading the configuration file at the given path
// and validating its contents. It exits the program if an error occurs.
func NewConfig(configFilePath string, log *slog.Logger) *Config {
	// Validate the config file
	if err := validateConfigFile(configFilePath); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Read the config file
	cfg, err := readConfigFile(configFilePath)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	return cfg
}

// validateConfigFile checks if the given config file exists and returns an error if it doesn't.
func validateConfigFile(configFilePath string) error {
	_, err := os.Stat(configFilePath)
	if err != nil {
		return err
	}
	return nil
}

// readConfigFile reads the configuration file at the given path and returns a Config struct.
func readConfigFile(configFilePath string) (*Config, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
