package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri, dbName string) (*mongo.Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Giảm xuống 5s cho nhanh
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }

    // THÊM ĐOẠN NÀY: Ép nó phải kiểm tra kết nối thật
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err // Nếu không có DB, nó sẽ trả về lỗi ngay tại đây
    }

    return client.Database(dbName), nil
}