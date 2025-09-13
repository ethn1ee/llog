package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Foo string `mapstructure:"foo"`
}

type configKey struct{}

func Init(cmd *cobra.Command, cfgFile string) error {
	ctx := cmd.Context()
	cfg := &Config{}

	setDefaults()
	viper.SetEnvPrefix("LLOG")
	viper.AutomaticEnv()

	if err := readConfig(cfgFile); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	cmd.SetContext(context.WithValue(ctx, configKey{}, cfg))

	return nil
}

func GetConfig(cmd *cobra.Command) (*Config, error) {
	v := cmd.Context().Value(configKey{})
	if v == nil {
		return nil, fmt.Errorf("config not found in context")
	}

	cfg, ok := v.(*Config)
	if !ok {
		return nil, fmt.Errorf("config in context is not of type *config.Config")
	}

	return cfg, nil
}

func readConfig(cfgFile string) error {
	if cfgFile != "" {
		if _, err := os.Stat(cfgFile); err != nil {
			return fmt.Errorf("provided config file does not exist")
		}
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to locate home directory: %w", err)
		}

		configPath := path.Join(home, ".config", "llog")

		viper.AddConfigPath(configPath)
		viper.SetConfigName("llog")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			slog.Warn("no config file found, using defaults")
			return nil
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	return nil
}

func setDefaults() {
	viper.SetDefault("foo", "bar")
}
