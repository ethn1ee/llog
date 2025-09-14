package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Entry *entryDB
}

func New(cfg *config.Config) (*DB, error) {
	dir := filepath.Dir(cfg.DBPath)
	if err := os.Mkdir(dir, 0755); err != nil && !os.IsExist(err) {
		return nil, fmt.Errorf("failed to create db directory: %w", err)
	}

	gormdb, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := gormdb.AutoMigrate(&models.Entry{}); err != nil {
		return nil, fmt.Errorf("failed to migrate Entry: %w", err)
	}

	db := &DB{
		Entry: &entryDB{gorm.G[models.Entry](gormdb)},
	}

	return db, nil
}
