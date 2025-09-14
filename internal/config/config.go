package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	ConfigDir string `mapstructure:"config_dir"`
	DBPath    string `mapstructure:"db_path"`
	LogPath   string `mapstructure:"log_path"`
}

type configKey struct{}

func Init(cmd *cobra.Command) error {
	ctx := cmd.Context()
	cfg := &Config{}

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

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	cmd.SetContext(context.WithValue(ctx, configKey{}, cfg))

	return nil
}

func FromCmd(cmd *cobra.Command) (*Config, error) {
	v := cmd.Context().Value(configKey{})
	if v == nil {
		return nil, errors.New("config not found in context")
	}

	cfg, ok := v.(*Config)
	if !ok {
		return nil, errors.New("config in context is not of type *config.Config")
	}

	return cfg, nil
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
