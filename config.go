package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultConfigDirectoryName = "commander"
	defaultConfigFileName      = "config"
	defaultConfigFileType      = "yml"
)

// Config is a struct for CLI configuration.
type Config struct {
	Name              *string
	FileName          *string
	FileType          *string
	Path              *string
	EnvironmentPrefix *string
}

// Init inits Config.
func (c *Config) Init() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("can't compose the default config file path: %v", err)
	}

	configName := defaultConfigDirectoryName
	if c.Name != nil {
		configName = *c.Name
	}

	configPath := path.Join(home, fmt.Sprintf(".%s", configName))
	if c.Path != nil {
		configPath = *c.Path
	}

	configFileName := defaultConfigFileName
	if c.FileName != nil {
		configFileName = *c.FileName
	}

	configFileType := defaultConfigFileType
	if c.FileType != nil {
		configFileType = *c.FileType
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)

	if c.EnvironmentPrefix != nil {
		viper.SetEnvPrefix(*c.EnvironmentPrefix)
	}

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
		if err != nil {
			return err
		}

		if err := viper.SafeWriteConfig(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) bindFlags(cmd *cobra.Command) error {
	var err error

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(flag.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(flag.Name, "-", "_"))
			err = viper.BindEnv(flag.Name, fmt.Sprintf("%s_%s", *c.EnvironmentPrefix, envVarSuffix))
			if err != nil {
				return
			}
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !flag.Changed && viper.IsSet(flag.Name) {
			val := viper.Get(flag.Name)
			err = cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
			if err != nil {
				return
			}
		}
	})

	return err
}
