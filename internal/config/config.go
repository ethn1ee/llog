package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigDir  string `mapstructure:"config_dir"`
	DBPath     string `mapstructure:"db_path"`
	LogPath    string `mapstructure:"log_path"`
	TimeLayout string `mapstructure:"time_layout"`
	DateLayout string `mapstructure:"date_layout"`
}

func Load(cfg *Config) error {
	configPath, err := createConfigPath()
	if err != nil {
		return fmt.Errorf("failed to create config path: %w", err)
	}

	setDefaults(configPath)

	viper.AddConfigPath(configPath)
	viper.SetConfigName("llog")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("LLOG")
	viper.AutomaticEnv()

	if err := readConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

func createConfigPath() (string, error) {
	// configDir, err := os.UserConfigDir()
	// if err != nil {
	// 	return "", fmt.Errorf("failed to locate user config directory: %w", err)
	// }
	configDir := "./.config"
	configPath := path.Join(configDir, "llog")

	if err := os.MkdirAll(configPath, 0o755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return "", fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	return configPath, nil
}

func readConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}
	}

	return nil
}