package handler

import (
	"chillhub/internal/module/media/service"
	"chillhub/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	service   *service.MediaService
	rawBucket string
}

func NewMediaHandler(
	s *service.MediaService,
) *MediaHandler {
	return &MediaHandler{
		service:   s,
	}
}

func (h *MediaHandler) InitUpload(c *gin.Context) {
	media, url, err := h.service.InitUpload(c.Request.Context())
	if err != nil {
		c.Error(err) //  giao toàn quyền cho global handler
		return
	}

	c.JSON(200, response.Envelope{
		Success: true,
		Data: gin.H{
			"id":         media.ID.Hex(),
			"upload_url": url,
		},
	})
}

