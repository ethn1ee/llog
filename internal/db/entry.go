package db

import (
	"context"
	"time"

	"github.com/ethn1ee/llog/internal/models"
	"gorm.io/gorm"
)

type entryDB struct {
	gorm.Interface[models.Entry]
}

func (db *entryDB) Add(ctx context.Context, entry *models.Entry) error {
	return db.Create(ctx, entry)
}

func (db *entryDB) GetAll(ctx context.Context) ([]models.Entry, error) {
	return db.Find(ctx)
}

func (db *entryDB) GetFrom(ctx context.Context, from time.Time) ([]models.Entry, error) {
	return db.Where("created_at >= ?", from).Find(ctx)
}

func (db *entryDB) GetTo(ctx context.Context, to time.Time) ([]models.Entry, error) {
	return db.Where("created_at <= ?", to).Find(ctx)
}

func (db *entryDB) GetRange(ctx context.Context, from time.Time, to time.Time) ([]models.Entry, error) {
	return db.Where("created_at >= ? AND created_at <= ?", from, to).Find(ctx)
}
