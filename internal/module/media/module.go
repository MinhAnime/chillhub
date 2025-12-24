package media

import (
	"chillhub/internal/module/media/handler"
	"chillhub/internal/module/media/repository"
	"chillhub/internal/module/media/service"
	minioshared "chillhub/internal/shared/minio"
)

// Module chứa handler để đăng ký route
const rawBucket = "raw-media-bucket" // mỗi module quản lý bucket riêng

type Module struct {
	Handler *handler.MediaHandler
}

func NewModule(repo repository.MediaRepository, minio *minioshared.Util) *Module {

	transcoder := service.NewTranscodingService()


	// Service nhận minio và bucket riêng
	svc := service.NewMediaService(repo, transcoder, minio, rawBucket)

	// Handler nhận service
	h := handler.NewMediaHandler(svc)

	return &Module{Handler: h}
}
