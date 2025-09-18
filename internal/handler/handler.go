package handler

import (
	"fmt"
	"log/slog"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/ethn1ee/llog/internal/view"
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

		cmdAttr := slog.Group(
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

		if err := _db.Load(cfg, ctx, db); err != nil {
			return fmt.Errorf("failed to load db: %w", err)
		}

		return nil
	}
}

func Add(cfg *config.Config, db *_db.DB, opts *AddOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		entry := &model.Entry{
			Body: args[0],
		}

		if err := db.Entry.Add(ctx, entry); err != nil {
			return fmt.Errorf(addEntryError, err)
		}

		slog.Info("added entry", slog.Any("entry", entry))

		return nil
	}
}

func Get(cfg *config.Config, db *_db.DB, opts *GetOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		var entries []model.Entry
		var err error

		from, to := opts.Time.fromTime, opts.Time.toTime

		if !from.IsZero() && !to.IsZero() {
			entries, err = db.Entry.GetRange(ctx, from, to)
		} else if !from.IsZero() {
			entries, err = db.Entry.GetFrom(ctx, from)
		} else if !to.IsZero() {
			entries, err = db.Entry.GetTo(ctx, to)
		} else {
			entries, err = db.Entry.GetAll(ctx)
		}

		if err != nil {
			return fmt.Errorf(getEntryError, err)
		}

		view.PrintEntries(cfg, entries)

		slog.Info("retrieved entries", slog.Int("count", len(entries)))

		return nil
	}
}

func Delete(cfg *config.Config, db *_db.DB, opts *DeleteOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		return nil
	}
}

func Summarize(cfg *config.Config, db *_db.DB, opts *SummarizeOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Println("summarize called")
		return nil
	}
}
