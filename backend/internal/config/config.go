package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io/fs"
	"os"
)

type database struct {
	DSN string
}

type twitter struct {
	ConsumerKey    string
	ConsumerSecret string
}

type Config struct {
	Port     uint
	Database database
	Twitter  twitter
}

func writeDefaultConfig(path string) error {
	content, err := toml.Marshal(&Config{
		Port:     8080,
		Database: database{DSN: ""},
		Twitter: twitter{
			ConsumerKey:    "",
			ConsumerSecret: "",
		},
	})
	if err != nil {
		return fmt.Errorf("couldn't generate default config: %w", err)
	}

	return os.WriteFile(path, content, fs.FileMode(0640))
}

func Read(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, writeDefaultConfig(path)
	}

	var config Config
	err = toml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
