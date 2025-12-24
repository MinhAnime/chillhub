package repository

import (
	"context"
	"time"

	"chillhub/internal/module/media/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MediaRepository interface {
	Insert(ctx context.Context, media *model.Media) error
	UpdateStatus(ctx context.Context, id string, status model.MediaStatus) error
}

type mongoRepo struct {
	col *mongo.Collection
}

// UpdateStatus implements [MediaRepository].
func (r *mongoRepo) UpdateStatus(ctx context.Context, id string, status model.MediaStatus) error {
	panic("unimplemented")
}

func NewMediaMongo(db *mongo.Database) MediaRepository {
	return &mongoRepo{
		col: db.Collection("media"),
	}
}

func (r *mongoRepo) Insert(ctx context.Context, media *model.Media) error {
	media.CreatedAt = time.Now().Unix()
	_, err := r.col.InsertOne(ctx, media)
	return err
}
