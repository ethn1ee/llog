package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/ethn1ee/llog/internal/view"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Get(cfg *config.Config, db *_db.DB, opts *GetOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return getById(cfg, db)(cmd, args)
		}
		return getWithOpts(cfg, db, opts)(cmd, args)
	}
}

func getById(cfg *config.Config, db *_db.DB) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		id, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid ID: %w", err)
		}

		if id < 1 || id > cfg.Internal.MaxEntryId {
			return fmt.Errorf("invalid ID: ID %d does not exist", id)
		}

		entry, err := db.Entry.GetById(ctx, id)
		if err != nil {
			return fmt.Errorf(getEntryError, err)
		}

		view.PrintEntry(cfg, entry)

		slog.Info("retrieved entry with ID", slog.Uint64("id", id))

		return nil
	}
}

func getWithOpts(cfg *config.Config, db *_db.DB, opts *GetOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		var entries []model.Entry
		var err error

		from, to := opts.Time.fromTime, opts.Time.toTime

		if !from.IsZero() || !to.IsZero() {
			entries, err = db.Entry.GetRange(ctx, from, to)
		} else {
			entries, err = db.Entry.GetAll(ctx)
		}

		if err != nil {
			return fmt.Errorf(getEntryError, err)
		}

		view.PrintEntries(cfg, entries)

		slog.Info("retrieved entries with options", slog.Int("count", len(entries)))

		return nil
	}
}

type GetOpts struct {
	Time  timeOpts
	Limit int
}

func (o *GetOpts) applyFlags(cmd *cobra.Command) {
	o.Time.applyFlags(cmd)
	cmd.Flags().IntVarP(&(o.Limit), "limit", "n", 10, "number of entries to return (default"+strconv.Itoa(10)+")")
}

func (o *GetOpts) validate(cfg *config.Config, cmd *cobra.Command, args []string) error {
	isIdSet := len(args) > 0
	isFlagSet := false

	cmd.Flags().Visit(func(f *pflag.Flag) {
		isFlagSet = true
	})

	if isIdSet && isFlagSet {
		return errors.New(flagIdMutexError)
	}

	return o.Time.validate(cfg, cmd, args)
}
