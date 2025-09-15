package config

import (
	"path"

	"github.com/spf13/viper"
)

func setDefaults(configPath string) {
	dbPath := path.Join(configPath, "db", "llog.db")
	logPath := path.Join(configPath, "log")

	viper.SetDefault("config_dir", configPath)
	viper.SetDefault("db_path", dbPath)
	viper.SetDefault("log_path", logPath)
	viper.SetDefault("time_layout", "2006-01-02 15:04")
	viper.SetDefault("date_layout", "2006-01-02")
}
