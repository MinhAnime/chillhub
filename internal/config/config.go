package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	MongoURI string
	MongoDB  string

	MinioEndpoint string
	MinioKey      string
	MinioSecret   string
	RawBucket     string
	HLSBucket     string
	UseSSL        bool
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Port:           os.Getenv("APP_PORT"),
		MongoURI:       os.Getenv("MONGO_URI"),
		MongoDB:        os.Getenv("MONGO_DB"),
		MinioEndpoint:  os.Getenv("MINIO_ENDPOINT"),
		MinioKey:       os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecret:   os.Getenv("MINIO_SECRET_KEY"),
		RawBucket:     os.Getenv("MINIO_BUCKET_RAW"),
		HLSBucket:     os.Getenv("MINIO_BUCKET_HLS"),
	}

	if cfg.Port == "" {
		log.Fatal("APP_PORT is required")
	}
	return cfg
}
