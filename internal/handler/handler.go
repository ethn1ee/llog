package handler

import (
	"fmt"
	"log/slog"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/spf13/cobra"
)

type HandlerFunc func(cmd *cobra.Command, args []string) error

const (
	addEntryError = "failed to add entry: %w"
	getEntryError = "failed to get entries: %w"
)

func Init(cfg *config.Config, db *_db.DB, lg *logger.Logger) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		if err := config.Load(cfg); err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if err := logger.Load(cfg, lg); err != nil {
			return fmt.Errorf("failed to load logger: %w", err)
		}

		slog.Info("command called",
			slog.Group(
				"command",
				slog.String("name", cmd.Name()),
				slog.Any("args", args),
			),
			slog.Any(
				"config",
				cfg,
			),
		)

		if err := _db.Load(cfg, ctx, db); err != nil {
			return fmt.Errorf("failed to load db: %w", err)
		}

		return nil
	}
}
