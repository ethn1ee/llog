package db

import (
	"context"
	"time"

	"github.com/ethn1ee/llog/internal/model"
	"gorm.io/gorm"
)

type entryDB struct {
	i gorm.Interface[model.Entry]
}

func (db *entryDB) Add(ctx context.Context, entry *model.Entry) error {
	return db.i.Create(ctx, entry)
}

func (db *entryDB) GetAll(ctx context.Context) ([]model.Entry, error) {
	return db.i.Find(ctx)
}

func (db *entryDB) GetFrom(ctx context.Context, from time.Time) ([]model.Entry, error) {
	return db.i.Where("created_at >= ?", from).Find(ctx)
}

func (db *entryDB) GetTo(ctx context.Context, to time.Time) ([]model.Entry, error) {
	return db.i.Where("created_at <= ?", to).Find(ctx)
}

func (db *entryDB) GetRange(ctx context.Context, from time.Time, to time.Time) ([]model.Entry, error) {
	return db.i.Where("created_at >= ? AND created_at <= ?", from, to).Find(ctx)
}

func (db *entryDB) GetLast(ctx context.Context) (model.Entry, error) {
	return db.i.Last(ctx)
}
