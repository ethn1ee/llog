package db

import (
	"context"
	"fmt"

	"github.com/ethn1ee/llog/internal/models"
	"gorm.io/gorm"
)

type entryDB struct {
	gorm.Interface[models.Entry]
}

func (db *entryDB) Add(ctx context.Context, entry *models.Entry) error {
	if err := db.Create(ctx, entry); err != nil {
		return fmt.Errorf("failed to add entry: %w", err)
	}

	return nil
}

func (db *entryDB) GetAll(ctx context.Context) ([]models.Entry, error) {
	return db.Find(ctx)
}
