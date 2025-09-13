/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/log"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "llog",
	Short: "Log your life",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := log.Init(); err != nil {
			return fmt.Errorf("failed to initialize logger: %w", err)
		}
		slog.Info("logger initialized")

		if err := config.Init(cmd, cfgFile); err != nil {
			return fmt.Errorf("failed to initialize config: %w", err)
		}
		slog.Info("config initialized")

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/llog.yaml)")
}
