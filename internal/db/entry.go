package db

import (
	"context"
	"errors"
	"time"

	"github.com/ethn1ee/llog/internal/model"
	"gorm.io/gorm"
)

type entryDB struct {
	i gorm.Interface[model.Entry]
}

func (db *entryDB) Count(ctx context.Context) (int64, error) {
	return db.i.Count(ctx, "id")
}

func (db *entryDB) Add(ctx context.Context, entry *model.Entry) error {
	return db.i.Create(ctx, entry)
}

func (db *entryDB) GetAll(ctx context.Context) ([]model.Entry, error) {
	return db.i.Find(ctx)
}

func (db *entryDB) GetRange(ctx context.Context, from time.Time, to time.Time) ([]model.Entry, error) {
	if from.IsZero() && to.IsZero() {
		return nil, errors.New("range unspecified")
	}
	if from.IsZero() {
		return db.i.Where("created_at <= ?", to).Find(ctx)
	}
	if to.IsZero() {
		return db.i.Where("created_at >= ?", from).Find(ctx)
	}
	return db.i.Where("created_at >= ? AND created_at <= ?", from, to).Find(ctx)
}

func (db *entryDB) GetLast(ctx context.Context) (model.Entry, error) {
	return db.i.Last(ctx)
}

func (db *entryDB) GetById(ctx context.Context, id uint64) (model.Entry, error) {
	match, err := db.i.Where("id = ?", id).Find(ctx)
	if err != nil {
		return model.Entry{}, err
	}

	return match[0], nil
}

func (db *entryDB) DeleteById(ctx context.Context, id uint64) error {
	_, err := db.i.Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
