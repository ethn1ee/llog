package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	entries gorm.Interface[models.Entry]
}

type dbKey struct{}

func Init(cmd *cobra.Command) error {
	ctx := cmd.Context()

	cfg, err := config.FromCmd(cmd)
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	dir := filepath.Dir(cfg.DBPath)
	if err := os.Mkdir(dir, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create db directory: %w", err)
	}

	gormdb, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}

	if err := gormdb.AutoMigrate(&models.Entry{}); err != nil {
		return fmt.Errorf("failed to migrate Entry: %w", err)
	}

	entries := gorm.G[models.Entry](gormdb)

	db := &DB{
		entries: entries,
	}

	cmd.SetContext(context.WithValue(ctx, dbKey{}, db))

	return nil
}

func FromCmd(cmd *cobra.Command) (*DB, error) {
	v := cmd.Context().Value(dbKey{})
	if v == nil {
		return nil, fmt.Errorf("db not found in context")
	}

	db, ok := v.(*DB)
	if !ok {
		return nil, fmt.Errorf("db in context is not of type *db.db")
	}

	return db, nil
}

func (db *DB) AddEntry(ctx context.Context, entry *models.Entry) error {
	if err := db.entries.Create(ctx, entry); err != nil {
		return fmt.Errorf("failed to add entry: %w", err)
	}

	return nil
}

func (db *DB) GetAllEntry(ctx context.Context) ([]models.Entry, error) {
	return db.entries.Find(ctx)
}
