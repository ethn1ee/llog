package config

import (
	"path"

	"github.com/spf13/viper"
)

func setDefaults(configPath string) {
	dbPath := path.Join(configPath, "db", "llog.db")
	logPath := path.Join(configPath, "log")

	viper.SetDefault("config_path", configPath)
	viper.SetDefault("db_path", dbPath)
	viper.SetDefault("log_path", logPath)
}
