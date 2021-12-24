package cmd

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	defaultConfigName     = "commander"
	defaultConfigFileName = "config"
	defaultConfigFileType = "yml"
)

type Config struct {
	Name *string
}

func (c *Config) Init() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("can't compose the default config file path: %v", err)
	}

	configName := defaultConfigName
	if c.Name != nil {
		configName = *c.Name
	}

	configPath := path.Join(home, fmt.Sprintf(".%s", configName))

	viper.AddConfigPath(configPath)
	viper.SetConfigName(defaultConfigFileName)
	viper.SetConfigType(defaultConfigFileType)
	viper.AutomaticEnv()

	err = c.ensureConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error when Fetching Configuration - %s", err)
	}

	return c
}

func (c *Config) ensureConfig(configPath string) error {
	configFilePath := path.Join(configPath, fmt.Sprintf("%s.%s", defaultConfigFileName, defaultConfigFileType))

	_, err := os.Stat(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path.Dir(configFilePath), os.ModePerm)
		if err := viper.SafeWriteConfig(); err != nil {
			return err
		}
	}

	return nil
}
