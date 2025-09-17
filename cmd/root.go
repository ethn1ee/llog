/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/spf13/cobra"
)

var (
	cfg     = &config.Config{}
	db      = &_db.DB{}
	lg      = &logger.Logger{}
	cmdAttr slog.Attr
)

var rootCmd = &cobra.Command{
	Use:   "llog",
	Short: "Life log",
	Long:  `Record your fleeting moments with llog.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cmdAttr = slog.Group(
			"command",
			slog.String("name", cmd.Name()),
			slog.Any("args", args),
		)

		if err := config.Load(cfg); err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if err := logger.Load(cfg, lg); err != nil {
			return fmt.Errorf("failed to load logger: %w", err)
		}

		slog.Info("command started", cmdAttr)
		slog.Info("using config", cmdAttr, slog.Any("config", cfg))

		if err := _db.Load(cfg, db); err != nil {
			return fmt.Errorf("failed to load db: %w", err)
		}

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("command completed", cmdAttr)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("command failed", cmdAttr, slog.Any("error", err))
		os.Exit(1)
	}
	if err := lg.Close(); err != nil {
		slog.Error("failed to close log file", cmdAttr, slog.Any("error", err))
	}
}
