package config

import (
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
}

func Init(cfgFile string) func() {
	return func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to locate home directory: %w", err))
			}
			cfgPath := path.Join(home, ".config", "llog")

			viper.AddConfigPath(cfgPath)
			viper.SetConfigType("yaml")
			viper.SetConfigName("llog")
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			slog.Info("Using config file", slog.String("file", viper.ConfigFileUsed()))
		}

		slog.Info("Config initialized")
	}
}
