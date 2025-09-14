package handler

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/models"
	"github.com/spf13/cobra"
)

type Handler struct {
	cfg *config.Config
	db  *db.DB
}

func New(cmd *cobra.Command) (*Handler, error) {
	cfg, err := config.FromCmd(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	db, err := db.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get db: %w", err)
	}

	return &Handler{
		cfg: cfg,
		db:  db,
	}, nil
}

func (h *Handler) AddEntry(cmd *cobra.Command, args []string) error {
	entry := &models.Entry{
		Body: strings.Join(args, " "),
	}

	if err := h.db.Entry.Add(cmd.Context(), entry); err != nil {
		return fmt.Errorf("failed to add entry: %w", err)
	}
	slog.Info("added entry", slog.Any("entry", entry))

	return nil
}

func (h *Handler) GetEntry(cmd *cobra.Command, args []string, opts *GetOpts) error {
	ctx := cmd.Context()

	var entries []models.Entry

	err := opts.validate()
	if err != nil {
		return fmt.Errorf("failed to validate options: %w", err)
	}

	if opts.From != "" && opts.To != "" {
		entries, err = h.db.Entry.GetRange(ctx, opts.fromTime, opts.toTime)
		if err != nil {
			return fmt.Errorf("failed to get entries: %w", err)
		}
	} else if opts.From != "" {
		entries, err = h.db.Entry.GetFrom(ctx, opts.fromTime)
		if err != nil {
			return fmt.Errorf("failed to get entries: %w", err)
		}
	} else if opts.To != "" {
		entries, err = h.db.Entry.GetTo(ctx, opts.toTime)
		if err != nil {
			return fmt.Errorf("failed to get entries: %w", err)
		}
	} else {
		entries, err = h.db.Entry.GetAll(ctx)
		if err != nil {
			return fmt.Errorf("failed to get entries: %w", err)
		}
	}

	for _, e := range entries {
		fmt.Println(e)
	}

	return nil
}

