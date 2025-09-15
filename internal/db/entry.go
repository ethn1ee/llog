package db

import (
	"context"
	"time"

	"github.com/ethn1ee/llog/internal/model"
	"gorm.io/gorm"
)

type entryDB struct {
	gorm.Interface[model.Entry]
}

func (db *entryDB) Add(ctx context.Context, entry *model.Entry) error {
	return db.Create(ctx, entry)
}

func (db *entryDB) GetAll(ctx context.Context) ([]model.Entry, error) {
	return db.Find(ctx)
}

func (db *entryDB) GetFrom(ctx context.Context, from time.Time) ([]model.Entry, error) {
	return db.Where("created_at >= ?", from).Find(ctx)
}

func (db *entryDB) GetTo(ctx context.Context, to time.Time) ([]model.Entry, error) {
	return db.Where("created_at <= ?", to).Find(ctx)
}

func (db *entryDB) GetRange(ctx context.Context, from time.Time, to time.Time) ([]model.Entry, error) {
	return db.Where("created_at >= ? AND created_at <= ?", from, to).Find(ctx)
}
