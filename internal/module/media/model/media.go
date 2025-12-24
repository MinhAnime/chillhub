// internal/module/media/model/media.go
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type MediaStatus string

const (
	StatusInit        MediaStatus = "INIT"
	StatusUploading   MediaStatus = "UPLOADING"
	StatusUploaded    MediaStatus = "UPLOADED"
	StatusTranscoding MediaStatus = "TRANSCODING"
	StatusReady       MediaStatus = "READY"
	StatusFailed      MediaStatus = "FAILED"
)

type Media struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  string             `bson:"title"`
	Status MediaStatus        `bson:"status"`

	Raw RawInfo `bson:"raw"`
	HLS HLSInfo `bson:"hls"`

	CreatedAt int64 `bson:"created_at"`
}

type RawInfo struct {
	Bucket string `bson:"bucket"`
	Object string `bson:"object"`
}

type HLSInfo struct {
	Bucket   string   `bson:"bucket"`
	Playlist string   `bson:"playlist"`
	Variants []string `bson:"variants"`
}
