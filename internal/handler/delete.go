package handler

import (
	"fmt"
	"strconv"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/spf13/cobra"
)

func Delete(cfg *config.Config, db *_db.DB, opts *DeleteOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		id, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid ID: %w", err)
		}

		if id < 1 || id > cfg.Internal.MaxEntryId {
			return fmt.Errorf("invalid ID: ID %d does not exist", id)
		}

		if err := db.Entry.DeleteById(ctx, id); err != nil {
			return fmt.Errorf("failed to delete entry: %w", err)
		}

		return nil
	}
}

type DeleteOpts struct {
	Interactive bool
}

func (o *DeleteOpts) applyFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&(o.Interactive), "interactive", "i", false, "select entries interactively")
}

func (o *DeleteOpts) validate(cfg *config.Config, cmd *cobra.Command, args []string) error {
	return nil
}
