package handler

import (
	"fmt"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/ethn1ee/llog/internal/model"
	"github.com/ethn1ee/llog/internal/view"
	"github.com/spf13/cobra"
)

func Add(cfg *config.Config, db *_db.DB, opts *AddOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		logger.LogCmdStart(cmd)
		defer logger.LogCmdComplete(cmd)

		ctx := cmd.Context()
		entries := make([]model.Entry, len(args))

		for i, arg := range args {
			entries[i] = model.Entry{
				Body: arg,
			}
		}

		if err := db.Entry.Add(ctx, entries); err != nil {
			return fmt.Errorf(dbAddEntryError, err)
		}

		view.PrintAdd(len(entries))

		return nil
	}
}

type AddOpts struct{}

func (o *AddOpts) applyFlags(cmd *cobra.Command) {}

func (o *AddOpts) validate(cfg *config.Config, args []string, flags []string) error { return nil }
