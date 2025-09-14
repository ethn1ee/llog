package handler

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

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
	db, err := db.FromCmd(cmd)
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
		Time: time.Now().Unix(),
		Body: strings.Join(args, " "),
	}

	if err := h.db.AddEntry(cmd.Context(), entry); err != nil {
		return fmt.Errorf("failed to add entry: %w", err)
	}
	slog.Info("added entry", slog.Any("entry", entry))

	return nil
}

func (h *Handler) GetAllEntry(cmd *cobra.Command, args []string) error {
	entries, err := h.db.GetAllEntry(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get all entries: %w", err)
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}

	return nil
}
