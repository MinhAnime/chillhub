package service

import (
	"context"
	"time"

	"chillhub/internal/module/media/model"
	"chillhub/internal/module/media/repository"
	minioshared "chillhub/internal/shared/minio"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MediaServiceInterface định nghĩa các hàm chính của MediaService
type MediaServiceInterface interface {
	// InitUpload tạo media record và trả về presigned URL
	InitUpload(ctx context.Context) (*model.Media, string, error)

	// Create lưu media vào repository và trigger transcoding
	Create(ctx context.Context, media *model.Media) error

	// PresignRawUpload trả về URL presigned PUT (nếu muốn expose riêng)
	PresignRawUpload(ctx context.Context, object string, expiry time.Duration) (string, error)
}


// MediaService implement MediaServiceInterface
type MediaService struct {
	repo       repository.MediaRepository
	transcoder *TranscodingService
	minio      *minioshared.Util
	rawBucket  string // thêm field này
}

// Constructor
func NewMediaService(
	repo repository.MediaRepository,
	transcoder *TranscodingService,
	minio *minioshared.Util,
	rawBucket string, // truyền bucket lúc init
) *MediaService {
	print("rawBucket trong service: ", rawBucket)
	return &MediaService{
		repo:       repo,
		transcoder: transcoder,
		minio:      minio,
		rawBucket:  rawBucket,
	}
}

// InitUpload tạo record media và trả về presigned PUT URL
func (s *MediaService) InitUpload(
	ctx context.Context,
) (*model.Media, string, error) {

	id := primitive.NewObjectID()
	object := "raw/" + id.Hex() // không cần extension, FE sẽ gắn extension

	print("Tại InitUpload: ",s.rawBucket)

	media := &model.Media{
		ID:     id,
		Status: model.StatusUploading,
		Raw: model.RawInfo{
			Bucket: s.rawBucket,
			Object: object,
		},
	}

	// Lưu vào repository
	if err := s.repo.Insert(ctx, media); err != nil {
		return nil, "", err
	}

	// Tạo presigned PUT URL
	url, err := s.minio.PresignPut(
		ctx,
		s.rawBucket,
		object,
		15*time.Minute,
	)

	if err := s.minio.EnsureBucket(context.Background(), s.rawBucket); err != nil {
		panic(err) // fail fast, config lỗi thì app không nên chạy
	}
	if err != nil {
		return nil, "", err
	}

	return media, url, nil
}

// Create lưu media và trigger transcoding
func (s *MediaService) Create(ctx context.Context, media *model.Media) error {
	if err := s.repo.Insert(ctx, media); err != nil {
		return err
	}

	// async transcoding
	go s.transcoder.Transcode(media)

	return nil
}

// PresignRawUpload trả về presigned PUT URL cho object
func (s *MediaService) PresignRawUpload(ctx context.Context, object string, expiry time.Duration) (string, error) {
	return s.minio.PresignPut(ctx, s.rawBucket, object, expiry)
}