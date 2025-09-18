package handler

import (
	"fmt"
	"log/slog"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/spf13/cobra"
)

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

type AddOpts struct{}

func (o *AddOpts) applyFlags(cmd *cobra.Command) {}

func (o *AddOpts) validate(cfg *config.Config, cmd *cobra.Command, args []string) error { return nil }
